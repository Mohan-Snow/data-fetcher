package repo

import "time"

type CoinMarketEntity struct {
	CoinName    string
	PriceUsd    string
	LastUpdated time.Time
}
