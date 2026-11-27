package middleware

import (
	"errors"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/protengplus/proteng-bff/utils/apiutil"
)

func RequireTrustedOrigin(allowed []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions:
			ctx.Next()
			return
		}

		origin := ctx.GetHeader("Origin")
		if origin != "" && !slices.Contains(allowed, origin) {
			apiutil.ApiResponseForbidden(ctx, errors.New("untrusted origin"))
			return
		}
		if origin == "" && hasAuthCookie(ctx) {
			apiutil.ApiResponseForbidden(ctx, errors.New("missing origin"))
			return
		}

		ctx.Next()
	}
}

func hasAuthCookie(ctx *gin.Context) bool {
	_, accessErr := ctx.Cookie("access_token")
	_, refreshErr := ctx.Cookie("refresh_token")
	return accessErr == nil || refreshErr == nil
}
