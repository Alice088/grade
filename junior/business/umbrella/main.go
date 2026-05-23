package main

import (
	"alice088/umbrella/internal/middleware"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var mainTime = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.ReadButtonDetector())

	r.GET("/time", func(ctx *gin.Context) {
		now := time.Now().UTC()
		left := now.UTC().Sub(mainTime)

		days := int(left.Hours()) / 24
		hours := int(left.Hours()) % 24
		minutes := int(left.Minutes()) % 60

		ctx.String(http.StatusOK, fmt.Sprintf("left since [Jun 1 2025]: %dD-%dH-%dM\n", days, hours, minutes))
	})

	r.Run(":8080")
}
