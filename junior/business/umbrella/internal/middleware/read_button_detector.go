package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ReadButtonDetector() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("User-Role") == "admin" {
			log.Println("red button user detected")
		}
		ctx.Next()
	}
}