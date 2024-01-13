package middleware

import (
	"os"
	"strings"

	"proteng-bff/utils"
	"proteng-bff/utils/apiutil"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
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

		sub, err := utils.ValidateToken(access_token, os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"))
		if err != nil {
			apiutil.ApiResponseUnauthorized(ctx, err)
			return
		}

		ctx.Set("userId", sub)
		ctx.Next()
	}
}
