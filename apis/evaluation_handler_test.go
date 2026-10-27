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

func TestEvaluationHandlers(t *testing.T) {
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
	router.GET("/evaluation-plugins", GetEvaluationPlugins)
	router.GET("/jobs/:id/evaluations", GetJobEvaluations)
	router.POST("/jobs/:id/evaluations", RerunEvaluation)
	router.GET("/evaluations/:id", GetEvaluation)

	send := func(method string, path string, body string) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, httptest.NewRequest(method, path, strings.NewReader(body)))
		if w.Code != http.StatusOK {
			t.Fatalf("%s %s: expected 200, got %d", method, path, w.Code)
		}
	}

	t.Run("plugins are forwarded", func(tt *testing.T) {
		send(http.MethodGet, "/evaluation-plugins", "")
		if got.Method != http.MethodGet || got.URL.Path != "/evaluation-plugins" {
			tt.Fatalf("unexpected forward: %s %s", got.Method, got.URL.String())
		}
	})

	t.Run("job evaluations force the caller's user_id", func(tt *testing.T) {
		send(http.MethodGet, "/jobs/j1/evaluations?user_id=someone-else", "")
		if got.URL.Path != "/jobs/j1/evaluations" || len(got.URL.Query()["user_id"]) != 1 || got.URL.Query().Get("user_id") != "owner" {
			tt.Fatalf("unexpected forward: %s", got.URL.String())
		}
	})

	t.Run("rerun forces the caller's user_id in the body", func(tt *testing.T) {
		send(http.MethodPost, "/jobs/j1/evaluations", `{"user_id":"someone-else","plugins":["mock"],"mutation_id":"m1"}`)
		if got.Method != http.MethodPost || got.URL.Path != "/jobs/j1/evaluations" {
			tt.Fatalf("unexpected forward: %s %s", got.Method, got.URL.Path)
		}
		plugins, _ := gotBody["plugins"].([]interface{})
		if gotBody["user_id"] != "owner" || len(plugins) != 1 || plugins[0] != "mock" || gotBody["mutation_id"] != "m1" {
			tt.Fatalf("unexpected body: %v", gotBody)
		}
	})

	t.Run("evaluation forces the caller's user_id", func(tt *testing.T) {
		send(http.MethodGet, "/evaluations/r1?user_id=someone-else", "")
		if got.URL.Path != "/evaluations/r1" || len(got.URL.Query()["user_id"]) != 1 || got.URL.Query().Get("user_id") != "owner" {
			tt.Fatalf("unexpected forward: %s", got.URL.String())
		}
	})
}
