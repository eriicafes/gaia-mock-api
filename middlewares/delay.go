package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func UseDelay(delay time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time.Sleep(delay)

		ctx.Next()
	}
}
