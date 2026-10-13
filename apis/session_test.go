package apis

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/protengplus/proteng-bff/configs"
	"github.com/protengplus/proteng-bff/internal/logger"
)

func TestMain(m *testing.M) {
	logger.InitZap()
	os.Exit(m.Run())
}

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
		{"success sets both cookies and hides tokens", http.StatusOK, `{"code":200,"data":{"id":"u1","access_token":"jwt-abc","refresh_token":"r-xyz"}}`, true},
		{"failure sets no cookies", http.StatusBadRequest, `{"code":400,"error":"invalid email or password"}`, false},
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
			wantCookies := []struct{ name, value, path string }{
				{accessTokenCookie, "jwt-abc", "/api/"},
				{refreshTokenCookie, "r-xyz", "/api/proteng-user-mgmt/auth"},
			}
			for _, want := range wantCookies {
				cookie := findCookie(w.Result().Cookies(), want.name)
				if !tc.wantCookie {
					if cookie != nil {
						t.Fatalf("unexpected cookie %+v", cookie)
					}
					continue
				}
				if cookie == nil || cookie.Value != want.value || !cookie.HttpOnly || !cookie.Secure || cookie.SameSite != http.SameSiteLaxMode || cookie.Path != want.path || cookie.MaxAge != 0 {
					t.Fatalf("unexpected %s cookie %+v", want.name, cookie)
				}
			}
			if !tc.wantCookie {
				return
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

	cases := []struct {
		name          string
		cookie        string
		status        int
		down          bool
		wantForwarded bool
	}{
		{"revokes the refresh token upstream", "r-abc", http.StatusOK, false, true},
		{"upstream error still logs out", "r-abc", http.StatusInternalServerError, false, true},
		{"upstream down still logs out", "r-abc", http.StatusOK, true, false},
		{"no refresh cookie skips upstream", "", http.StatusOK, false, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := map[string]string{}
			fakeUserMgmt(t, tc.status, `{"code":200}`, &got)
			if tc.down {
				configs.Config.UserMgmtUrl = "http://127.0.0.1:0"
			}
			router := gin.New()
			router.POST("/logout", Logout)
			req := httptest.NewRequest(http.MethodPost, "/logout", nil)
			req.AddCookie(&http.Cookie{Name: accessTokenCookie, Value: "jwt-abc"})
			if tc.cookie != "" {
				req.AddCookie(&http.Cookie{Name: refreshTokenCookie, Value: tc.cookie})
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
			}
			if forwarded := got["refresh_token"] == "r-abc"; forwarded != tc.wantForwarded {
				t.Errorf("forwarded = %v, want %v", forwarded, tc.wantForwarded)
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
		})
	}
}

func TestRefreshSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
	configs.Config.CookieSecure = true
	configs.Config.CookiePathPrefix = "/api"

	cases := []struct {
		name          string
		cookie        string
		status        int
		body          string
		want          int
		wantSet       bool
		wantCleared   bool
		wantForwarded bool
	}{
		{"success sets new token cookies", "r-old", http.StatusOK, `{"code":200,"data":{"access_token":"jwt-new","refresh_token":"r-new"}}`, http.StatusOK, true, false, true},
		{"expired session clears cookies", "r-old", http.StatusUnauthorized, `{"code":401,"error":"refresh token is invalid or expired"}`, http.StatusUnauthorized, false, true, true},
		{"concurrent refresh keeps cookies", "r-old", http.StatusConflict, `{"code":409,"error":"refresh token was already used"}`, http.StatusConflict, false, false, true},
		{"missing cookie is rejected without calling user-mgmt", "", http.StatusOK, `{}`, http.StatusUnauthorized, false, true, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := map[string]string{}
			fakeUserMgmt(t, tc.status, tc.body, &got)
			router := gin.New()
			router.POST("/refresh", RefreshSession)
			req := httptest.NewRequest(http.MethodPost, "/refresh", nil)
			if tc.cookie != "" {
				req.AddCookie(&http.Cookie{Name: refreshTokenCookie, Value: tc.cookie})
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.want {
				t.Fatalf("status = %d, want %d", w.Code, tc.want)
			}
			if forwarded := got["refresh_token"] == tc.cookie && tc.cookie != ""; forwarded != tc.wantForwarded {
				t.Errorf("forwarded = %v, want %v (got %v)", forwarded, tc.wantForwarded, got)
			}

			cookies := w.Result().Cookies()
			switch {
			case tc.wantSet:
				access, refresh := findCookie(cookies, accessTokenCookie), findCookie(cookies, refreshTokenCookie)
				if access == nil || access.Value != "jwt-new" || refresh == nil || refresh.Value != "r-new" || refresh.Path != "/api/proteng-user-mgmt/auth" {
					t.Fatalf("cookies not rotated: %+v", cookies)
				}
				if strings.Contains(w.Body.String(), "jwt-new") || strings.Contains(w.Body.String(), "r-new") {
					t.Fatalf("token leaked: %s", w.Body.String())
				}
			case tc.wantCleared:
				for _, name := range []string{accessTokenCookie, refreshTokenCookie} {
					if cookie := findCookie(cookies, name); cookie == nil || cookie.MaxAge >= 0 {
						t.Errorf("%s not cleared: %+v", name, cookie)
					}
				}
			default:
				if len(cookies) != 0 {
					t.Fatalf("unexpected cookies %+v", cookies)
				}
			}
		})
	}
}
