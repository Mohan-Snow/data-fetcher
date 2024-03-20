package repo

import (
	"data-fetcher/internal"
	"data-fetcher/internal/model/repo"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
)

type postgresRepo struct {
	db *sql.DB
}

func NewCurrencyRepository(db *sql.DB) internal.PostgresRepo {
	return &postgresRepo{
		db: db,
	}
}

func (r postgresRepo) Save(entity *repo.CoinMarketEntity) error {
	toSql, args, err := sq.Insert("currency").
		Columns("coin_name", "price_usd", "last_updated").
		Values(entity.CoinName, entity.PriceUsd, entity.LastUpdated).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(toSql, args...)

	if err != nil {
		return err
	}
	return nil
}

func (r postgresRepo) GetAll() ([]*repo.CoinMarketEntity, error) {
	var coinMarketEntities []*repo.CoinMarketEntity
	query := sq.Select(
		"coin_name", "price_usd", "last_updated",
	).From(
		"currency",
	)
	stmt, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(stmt, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		entity := new(repo.CoinMarketEntity)
		if err := rows.Scan(&entity.CoinName, &entity.PriceUsd, &entity.LastUpdated); err != nil {
			return nil, err
		}
		coinMarketEntities = append(coinMarketEntities, entity)
	}
	return coinMarketEntities, nil
}
