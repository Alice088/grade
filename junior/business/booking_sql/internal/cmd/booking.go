package main

import (
	"alice088/booking_sql/internal/config"
	"context"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env: " + err.Error())
	}
}

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Failed to parse .env: " + err.Error())
	}

	conn, err := pgx.Connect(context.Background(), cfg.DB.URL)
	if err != nil {
		log.Fatal("Failed to connect DB " + err.Error())
	}
}
