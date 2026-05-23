package main

import (
	"alice088/booking_sql/internal/booking"
	"alice088/booking_sql/internal/postgres"

	"alice088/booking_sql/internal/config"
	"context"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env: %s", err)
	}
}

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse .env: %s", err)
	}

	conn, err := pgx.Connect(context.Background(), cfg.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect DB: %s", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping DB: %s", err)
	}

	r := gin.New()
	r.Use(gin.Logger())

	v1 := r.Group("api/v1")

	bookingRepo := postgres.NewBookingRepository(conn)
	bookingService := booking.NewService(bookingRepo)
	bookingHandler := booking.NewHandler(bookingService)
	
	bookingRoute := v1.Group("/booking")
	{
		bookingRoute.POST("/", bookingHandler.Book)
	}

}
