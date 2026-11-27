package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

func TestCheckResetToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cases := []struct {
		name   string
		status int
		body   string
	}{
		{"valid token", http.StatusOK, `{"code":200,"message":"Token is valid"}`},
		{"invalid token", http.StatusBadRequest, `{"code":400,"error":"invalid token"}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var gotMethod, gotPath string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath = r.Method, r.URL.Path
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.body))
			}))
			defer server.Close()
			configs.Config.UserMgmtUrl = server.URL

			router := gin.New()
			router.GET("/auth/resetpassword/:resetToken", CheckResetToken)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/auth/resetpassword/abc", nil))

			if w.Code != tc.status {
				t.Errorf("status = %d, want %d", w.Code, tc.status)
			}
			if gotMethod != http.MethodGet || gotPath != "/auth/resetpassword/abc" {
				t.Errorf("forwarded %s %s", gotMethod, gotPath)
			}
			if cookies := w.Result().Cookies(); len(cookies) != 0 {
				t.Errorf("cookies = %v, want none", cookies)
			}
		})
	}
}
