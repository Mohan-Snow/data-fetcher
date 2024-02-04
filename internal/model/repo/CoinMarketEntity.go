package repo

import "time"

type CmEntity struct {
	CoinName    string
	PriceUsd    string
	LastUpdated time.Time
}
