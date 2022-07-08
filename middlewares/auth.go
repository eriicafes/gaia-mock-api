package middlewares

import (
	"net/http"
	"strings"

	"github.com/eriicafes/go-api-starter/modules/auth"
	"github.com/eriicafes/go-api-starter/response"
	"github.com/gin-gonic/gin"
)

func UseAuth(authService auth.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.New(ctx)

		bearer := ctx.GetHeader("Authorization")

		if bearer == "" {
			res.SetStatus(http.StatusUnauthorized).ErrJSON()
			return
		}

		parts := strings.Split(bearer, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			res.SetStatus(http.StatusUnauthorized).SetMessage("malformed token").ErrJSON()
			return
		}

		user, err := authService.Profile(parts[1])

		if err != nil {
			res.SetStatus(http.StatusUnauthorized).ErrJSON()
			return
		}

		ctx.Set("auth", user)

		ctx.Next()
	}
}
