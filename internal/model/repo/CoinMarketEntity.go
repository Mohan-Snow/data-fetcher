package repo

import (
	"fmt"
	"time"
)

type CoinMarketEntity struct {
	CoinName    string
	PriceUsd    string
	LastUpdated time.Time
}

func (c CoinMarketEntity) String() string {
	return fmt.Sprintf("CoinName: %s, PriceUsd: %s, LastUpdated: %s", c.CoinName, c.PriceUsd, c.LastUpdated)
}
