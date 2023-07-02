package middleware

import (
	"Final/pkg/web"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			ctx.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, 401, errors.New("invalid token"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
