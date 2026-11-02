package apis

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
)

func fakeUserMgmt(t *testing.T, status int, body string, got *map[string]string) string {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got != nil {
			json.NewDecoder(r.Body).Decode(got)
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	configs.Config.UserMgmtUrl = server.URL
	return server.URL
}

func findCookie(cookies []*http.Cookie, name string) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func TestForwardIssuingSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
	configs.Config.CookieSecure = true
	configs.Config.CookiePathPrefix = "/api"

	cases := []struct {
		name       string
		status     int
		body       string
		wantCookie bool
	}{
		{"success sets cookie and hides tokens", http.StatusOK, `{"code":200,"data":{"id":"u1","access_token":"jwt-abc","refresh_token":"r-xyz"}}`, true},
		{"failure sets no cookie", http.StatusBadRequest, `{"code":400,"error":"invalid email or password"}`, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			url := fakeUserMgmt(t, tc.status, tc.body, nil)
			router := gin.New()
			router.POST("/login", ForwardIssuingSession(url))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"a@b.com"}`)))

			if w.Code != tc.status {
				t.Fatalf("status = %d, want %d", w.Code, tc.status)
			}
			cookie := findCookie(w.Result().Cookies(), accessTokenCookie)
			if !tc.wantCookie {
				if cookie != nil {
					t.Fatalf("unexpected cookie %+v", cookie)
				}
				return
			}
			if cookie == nil || cookie.Value != "jwt-abc" || !cookie.HttpOnly || !cookie.Secure || cookie.SameSite != http.SameSiteLaxMode || cookie.Path != "/api/" {
				t.Fatalf("unexpected cookie %+v", cookie)
			}
			var body struct {
				Data map[string]interface{} `json:"data"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			for _, field := range []string{"access_token", "refresh_token"} {
				if _, ok := body.Data[field]; ok {
					t.Fatalf("%s leaked: %s", field, w.Body.String())
				}
			}
			if body.Data["id"] != "u1" {
				t.Fatalf("user data lost: %s", w.Body.String())
			}
		})
	}
}

func TestLogout(t *testing.T) {
	gin.SetMode(gin.TestMode)
	configs.Config.CookiePathPrefix = "/api"

	router := gin.New()
	router.POST("/logout", Logout)
	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.AddCookie(&http.Cookie{Name: accessTokenCookie, Value: "jwt-abc"})
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}
	wantPaths := map[string]string{
		accessTokenCookie:  "/api/",
		refreshTokenCookie: "/api/proteng-user-mgmt/auth",
	}
	for name, path := range wantPaths {
		cookie := findCookie(w.Result().Cookies(), name)
		if cookie == nil || cookie.MaxAge >= 0 || cookie.Value != "" || cookie.Path != path {
			t.Errorf("%s not cleared: %+v", name, cookie)
		}
	}
}
