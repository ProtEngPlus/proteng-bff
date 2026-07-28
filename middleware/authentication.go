package middleware

import (
	"fmt"
	"slices"
	"strings"

	"github.com/protengplus/proteng-bff/configs"
	"github.com/protengplus/proteng-bff/utils"
	"github.com/protengplus/proteng-bff/utils/apiutil"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			apiutil.ApiResponseUnauthorized(ctx, err, "You are not logged in")
			return
		}

		sub, role, err := utils.ValidateToken(access_token, configs.Config.AccessTokenPublicKey)
		if err != nil {
			apiutil.ApiResponseUnauthorized(ctx, err)
			return
		}

		ctx.Set("userId", sub)
		ctx.Set("role", role)
		ctx.Next()
	}
}

func Authorize(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			apiutil.ApiResponseUnauthorized(ctx, fmt.Errorf("User not authenticated"))
			return
		}

		if !slices.Contains(roles, role.(string)) {
			apiutil.ApiResponseForbidden(ctx, fmt.Errorf("Insufficient permissions"))
			return
		}

		ctx.Next()
	}
}
