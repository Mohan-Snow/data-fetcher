package internal

import "data-fetcher/internal/model/repo"

type PostgresRepo interface {
	Save(entity *repo.CmEntity) error
	GetById(id int) (*repo.CmEntity, error)
	GetAll() ([]*repo.CmEntity, error)
}

type IFetchDataService interface {
	FetchDataAndSave(symbol string, token string) (string, error)
}
