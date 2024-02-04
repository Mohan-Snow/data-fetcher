package coinmarket

import (
	"data-fetcher/internal/model/api/coinmarket"
	"data-fetcher/internal/model/repo"
	"errors"
	"fmt"
)

func MapDtoToEntity(dto *coinmarket.CmDto) (*repo.CmEntity, error) {
	if dto.Data.BTC != nil {
		btc := dto.Data.BTC[0]
		entity := repo.CmEntity{
			CoinName:    btc.Name,
			PriceUsd:    fmt.Sprintf("%f", btc.Quote.USD.Price),
			LastUpdated: btc.LastUpdated,
		}
		return &entity, nil
	} else if dto.Data.ETH != nil {
		eth := dto.Data.ETH[0]
		entity := repo.CmEntity{
			CoinName:    eth.Name,
			PriceUsd:    fmt.Sprintf("%f", eth.Quote.USD.Price),
			LastUpdated: eth.LastUpdated,
		}
		return &entity, nil
	} else {
		return nil, errors.New("data was empty. Nothing is to map")
	}
}
