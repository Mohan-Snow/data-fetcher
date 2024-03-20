package internal

import "data-fetcher/internal/model/repo"

type PostgresRepo interface {
	Save(entity *repo.CoinMarketEntity) error
	GetAll() ([]*repo.CoinMarketEntity, error)
}

type IFetchDataService interface {
	FetchDataAndSave(symbol string, token string) (string, error)
	FetchAllData() ([]*repo.CoinMarketEntity, error)
}
