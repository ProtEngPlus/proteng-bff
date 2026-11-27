package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRequireTrustedOrigin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(RequireTrustedOrigin([]string{"https://www.protengplus.com"}))
	router.Any("/x", func(c *gin.Context) { c.Status(http.StatusOK) })

	cases := []struct {
		name   string
		method string
		origin string
		cookie string
		want   int
	}{
		{"get without origin", http.MethodGet, "", "access_token", http.StatusOK},
		{"post from allowed origin", http.MethodPost, "https://www.protengplus.com", "access_token", http.StatusOK},
		{"post from other origin", http.MethodPost, "https://evil.example", "", http.StatusForbidden},
		{"post from other origin with cookie", http.MethodPost, "https://evil.example", "access_token", http.StatusForbidden},
		{"post without origin or cookie", http.MethodPost, "", "", http.StatusOK},
		{"post without origin but with access cookie", http.MethodPost, "", "access_token", http.StatusForbidden},
		{"delete without origin but with refresh cookie", http.MethodDelete, "", "refresh_token", http.StatusForbidden},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/x", nil)
			if tc.origin != "" {
				req.Header.Set("Origin", tc.origin)
			}
			if tc.cookie != "" {
				req.AddCookie(&http.Cookie{Name: tc.cookie, Value: "token"})
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Fatalf("status = %d, want %d", w.Code, tc.want)
			}
		})
	}
}
