package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/protengplus/proteng-bff/configs"
	"github.com/protengplus/proteng-bff/models"
	"github.com/protengplus/proteng-bff/utils/apiutil"
)

const (
	accessTokenCookie  = "access_token"
	refreshTokenCookie = "refresh_token"
	accessCookiePath   = "/"
	refreshCookiePath  = "/proteng-user-mgmt/auth"
)

type sessionResponse = models.HttpResponseWithData[map[string]interface{}]

var userMgmtClient = &http.Client{Timeout: 10 * time.Second}

func setAuthCookie(c *gin.Context, name, value, path string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     configs.Config.CookiePathPrefix + path,
		HttpOnly: true,
		Secure:   configs.Config.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearAuthCookie(c *gin.Context, name, path string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Path:     configs.Config.CookiePathPrefix + path,
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   configs.Config.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})
}

func clearAuthCookies(c *gin.Context) {
	clearAuthCookie(c, accessTokenCookie, accessCookiePath)
	clearAuthCookie(c, refreshTokenCookie, refreshCookiePath)
}

func moveTokensToCookies(c *gin.Context, data map[string]interface{}) {
	if token, ok := data["access_token"].(string); ok && token != "" {
		setAuthCookie(c, accessTokenCookie, token, accessCookiePath)
	}
	if token, ok := data["refresh_token"].(string); ok && token != "" {
		setAuthCookie(c, refreshTokenCookie, token, refreshCookiePath)
	}
	delete(data, "access_token")
	delete(data, "refresh_token")
}

func callUserMgmt(c *gin.Context, method, url string, body []byte) (int, sessionResponse, error) {
	var out sessionResponse
	req, err := http.NewRequestWithContext(c.Request.Context(), method, url, bytes.NewReader(body))
	if err != nil {
		return 0, out, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := userMgmtClient.Do(req)
	if err != nil {
		return 0, out, fmt.Errorf("call user-mgmt: %w", err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return 0, out, fmt.Errorf("decode user-mgmt response: %w", err)
	}
	return resp.StatusCode, out, nil
}

// ForwardIssuingSession forwards a login-like request and moves the issued token into an httpOnly cookie
func ForwardIssuingSession(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			apiutil.ApiResponseErrorBadRequest(c, err)
			return
		}
		status, out, err := callUserMgmt(c, http.MethodPost, url, body)
		if err != nil {
			apiutil.ApiResponseInternalServerError(c, err)
			return
		}
		if status == http.StatusOK && out.Data != nil {
			moveTokensToCookies(c, out.Data)
		}
		c.JSON(status, out)
	}
}

func forwardRefresh(c *gin.Context, url string) {
	refreshToken, err := c.Cookie(refreshTokenCookie)
	if err != nil || refreshToken == "" {
		clearAuthCookies(c)
		apiutil.ApiResponseUnauthorized(c, errors.New("missing refresh token"), "error: session expired")
		return
	}
	body, err := json.Marshal(map[string]string{"refresh_token": refreshToken})
	if err != nil {
		apiutil.ApiResponseInternalServerError(c, err)
		return
	}
	status, out, err := callUserMgmt(c, http.MethodPost, url, body)
	if err != nil {
		apiutil.ApiResponseInternalServerError(c, err)
		return
	}
	switch status {
	case http.StatusOK:
		if out.Data != nil {
			moveTokensToCookies(c, out.Data)
		}
	case http.StatusUnauthorized:
		clearAuthCookies(c)
	}
	c.JSON(status, out)
}
