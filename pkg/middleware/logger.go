package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		verb := ctx.Request.Method
		path := ctx.Request.RequestURI

		ctx.Next()
		var size int
		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}
		fmt.Printf("time: %v\npath: localhost:8080%s\nverb: %s\nsize: %d\n", t, path, verb, size)
	}
}
