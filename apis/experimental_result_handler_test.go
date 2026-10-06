package apis

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

func TestExperimentalResultHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var got *http.Request
	var gotBody map[string]interface{}
	conductor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r
		gotBody = nil
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &gotBody)
		w.Write([]byte(`{"code":200}`))
	}))
	defer conductor.Close()
	configs.Config.ConductorUrl = conductor.URL

	router := gin.New()
	router.Use(func(c *gin.Context) { c.Set("userId", "owner") })
	router.GET("/mutations/experimental-results", GetAllExperimentalResults)
	router.PUT("/mutations/experimental-results/by-mutation-result/:mutationResultId", UpsertExperimentalResult)
	router.DELETE("/mutations/experimental-results/by-mutation-result/:mutationResultId", DeleteExperimentalResult)

	send := func(method string, path string, body string) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, httptest.NewRequest(method, path, strings.NewReader(body)))
		if w.Code != http.StatusOK {
			t.Fatalf("%s %s: expected 200, got %d", method, path, w.Code)
		}
	}

	t.Run("get forces the caller's user_id", func(tt *testing.T) {
		send(http.MethodGet, "/mutations/experimental-results?job_id=j&user_id=someone-else", "")
		if got.URL.Path != "/mutations/experimental-results" || got.URL.Query().Get("job_id") != "j" || got.URL.Query()["user_id"][0] != "owner" || len(got.URL.Query()["user_id"]) != 1 {
			tt.Fatalf("unexpected forward: %s", got.URL.String())
		}
	})

	t.Run("put forces the caller's user_id in the body", func(tt *testing.T) {
		send(http.MethodPut, "/mutations/experimental-results/by-mutation-result/abc", `{"user_id":"someone-else","actual_assay_score":1.5}`)
		if got.Method != http.MethodPut || got.URL.Path != "/mutations/experimental-results/by-mutation-result/abc" {
			tt.Fatalf("unexpected forward: %s %s", got.Method, got.URL.Path)
		}
		if gotBody["user_id"] != "owner" || gotBody["actual_assay_score"] != 1.5 {
			tt.Fatalf("unexpected body: %v", gotBody)
		}
	})

	t.Run("delete forces the caller's user_id", func(tt *testing.T) {
		send(http.MethodDelete, "/mutations/experimental-results/by-mutation-result/abc?user_id=someone-else", "")
		if got.Method != http.MethodDelete || got.URL.Path != "/mutations/experimental-results/by-mutation-result/abc" || got.URL.Query().Get("user_id") != "owner" {
			tt.Fatalf("unexpected forward: %s %s", got.Method, got.URL.String())
		}
	})
}
