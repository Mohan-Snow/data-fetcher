package coinmarketcap

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type CoinMarketConfig struct {
	CoinMarketUrl       string `env:"COIN_MARKET_CAP"`
	CoinMarketHeaderKey string `env:"COIN_MARKET_HEADER_KEY"`
}

func NewConfig() (*CoinMarketConfig, error) {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Print("No .env file found")
	}
	cfg := CoinMarketConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Printf("%+v\n", cfg)
	return &cfg, nil
}
