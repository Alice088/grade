package main

import (
	"alice088/umbrella/internal/handlers"
	"alice088/umbrella/internal/middleware"
	"alice088/umbrella/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	timeHandler := handlers.NewTime(
		service.NewTime(
			time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
		),
	)

	timeRoute := r.Group("/time")
	timeRoute.GET("/since", timeHandler.Since()).Use(middleware.ReadButtonDetector())

	r.Run(":8080")
}
