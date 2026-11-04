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

func TestAssayTypeHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var got *http.Request
	var gotBody map[string]interface{}
	conductor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r
		gotBody = nil
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &gotBody)
		if strings.HasSuffix(r.URL.Path, "/export") {
			w.Header().Set("Content-Disposition", `attachment; filename="a.csv"`)
			w.Header().Set("X-Row-Count", "2")
			w.Header().Set("Content-Type", "text/csv")
			w.Write([]byte("MKV,1\nAAA,2\n"))
			return
		}
		w.Write([]byte(`{"code":200}`))
	}))
	defer conductor.Close()
	configs.Config.ConductorUrl = conductor.URL

	router := gin.New()
	router.Use(func(c *gin.Context) { c.Set("userId", "owner") })
	router.GET("/assay-types", GetAllAssayTypes)
	router.POST("/assay-types", CreateAssayType)
	router.PUT("/assay-types/:id", UpdateAssayType)
	router.DELETE("/assay-types/:id", DeleteAssayType)
	router.GET("/assay-types/:id/library/jobs", GetLibraryJobs)
	router.GET("/assay-types/:id/library", GetLibraryResults)
	router.GET("/assay-types/:id/library/export", ExportLibrary)
	router.GET("/mutations/experimental-results/export", ExportExperimentalResults)

	send := func(method string, path string, body string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, httptest.NewRequest(method, path, strings.NewReader(body)))
		if w.Code != http.StatusOK {
			t.Fatalf("%s %s: expected 200, got %d", method, path, w.Code)
		}
		return w
	}

	t.Run("get paths force the caller's user_id", func(tt *testing.T) {
		for _, path := range []string{"/assay-types", "/assay-types/a1/library/jobs?q=x", "/assay-types/a1/library?job_id=j"} {
			send(http.MethodGet, path+map[bool]string{true: "&", false: "?"}[strings.Contains(path, "?")]+"user_id=someone-else", "")
			ids := got.URL.Query()["user_id"]
			if len(ids) != 1 || ids[0] != "owner" {
				tt.Fatalf("%s: unexpected user_id %v", path, ids)
			}
		}
		if got.URL.Path != "/assay-types/a1/library" || got.URL.Query().Get("job_id") != "j" {
			tt.Fatalf("unexpected forward: %s", got.URL.String())
		}
	})

	t.Run("delete forces the caller's user_id", func(tt *testing.T) {
		send(http.MethodDelete, "/assay-types/a1?user_id=someone-else", "")
		if got.Method != http.MethodDelete || got.URL.Path != "/assay-types/a1" || got.URL.Query().Get("user_id") != "owner" {
			tt.Fatalf("unexpected forward: %s %s", got.Method, got.URL.String())
		}
	})

	t.Run("post and put force the caller's user_id in the body", func(tt *testing.T) {
		send(http.MethodPost, "/assay-types", `{"user_id":"someone-else","name":"Yield"}`)
		if got.URL.Path != "/assay-types" || gotBody["user_id"] != "owner" || gotBody["name"] != "Yield" {
			tt.Fatalf("unexpected post: %s %v", got.URL.Path, gotBody)
		}
		send(http.MethodPut, "/assay-types/a1", `{"user_id":"someone-else","name":"Yield2"}`)
		if got.Method != http.MethodPut || got.URL.Path != "/assay-types/a1" || gotBody["user_id"] != "owner" {
			tt.Fatalf("unexpected put: %s %v", got.URL.Path, gotBody)
		}
	})

	t.Run("export forwards the file and its headers", func(tt *testing.T) {
		w := send(http.MethodGet, "/assay-types/a1/library/export?ids=1,2&user_id=someone-else", "")
		if got.URL.Query().Get("user_id") != "owner" || got.URL.Query().Get("ids") != "1,2" {
			tt.Fatalf("unexpected forward: %s", got.URL.String())
		}
		if w.Body.String() != "MKV,1\nAAA,2\n" || w.Header().Get("X-Row-Count") != "2" || !strings.Contains(w.Header().Get("Content-Disposition"), "a.csv") {
			tt.Fatalf("unexpected response: %q %v", w.Body.String(), w.Header())
		}

		send(http.MethodGet, "/mutations/experimental-results/export?job_id=j&user_id=someone-else", "")
		if got.URL.Path != "/mutations/experimental-results/export" || got.URL.Query().Get("user_id") != "owner" || got.URL.Query().Get("job_id") != "j" {
			tt.Fatalf("unexpected forward: %s", got.URL.String())
		}
	})
}
