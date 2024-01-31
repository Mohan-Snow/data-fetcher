package postgres

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type DbPostgresConfig struct {
	DbPort     string `env:"DB_PORT"`
	DbHost     string `env:"DB_HOST"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
}

func NewConfig() (*DbPostgresConfig, error) {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Print("No .env file found")
	}
	cfg := DbPostgresConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Printf("%+v\n", cfg)
	return &cfg, nil
}
