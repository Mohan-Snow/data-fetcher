package bitgo

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type ApiBitgoConfig struct {
	BitgoUrl string `env:"BITGO_URL"`
}

func NewConfig() (*ApiBitgoConfig, error) {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Print("No .env file found")
	}
	cfg := ApiBitgoConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Printf("%+v\n", cfg)
	return &cfg, nil
}
