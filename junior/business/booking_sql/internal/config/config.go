package config

type Config struct {
	DB DB
}

type DB struct {
	URL string `env:"DB_URL,required"`
}
