package repo

import (
	"data-fetcher/internal"
	"data-fetcher/internal/model/repo"
	"database/sql"
)

type postgresRepo struct {
	db *sql.DB
}

func NewBitGoRepository(db *sql.DB) internal.PostgresRepo {
	return &postgresRepo{
		db: db,
	}
}

func (r postgresRepo) Save(entity *repo.CmEntity) error {
	_, err := r.db.Exec("INSERT INTO currency(COIN_NAME, PRICE_USD, LAST_UPDATED) VALUES ($1,$2,$3)",
		entity.CoinName, entity.PriceUsd, entity.LastUpdated)
	if err != nil {
		return err
	}
	return nil
}

func (r postgresRepo) GetById(id int) (*repo.CmEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (r postgresRepo) GetAll() ([]*repo.CmEntity, error) {
	//TODO implement me
	panic("implement me")
}
