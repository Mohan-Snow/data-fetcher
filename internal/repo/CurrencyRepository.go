package repo

import (
	"data-fetcher/internal"
	"data-fetcher/internal/model/repo"
	"database/sql"
	"errors"
	"fmt"
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
	fmt.Printf("Try to save new coin: %s\n", entity)
	existingCoin, err := r.findCoin(entity.CoinName)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		// Coin does not exist, perform insert
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
		fmt.Println("New coin inserted successfully.")
		return nil
	case err != nil:
		fmt.Println("Error occurred while checking for existing coin.")
		return err
	default:
		// Coin exists, perform update
		_, err := sq.Update(
			"currency",
		).Set(
			"price_usd", existingCoin.PriceUsd,
		).Set(
			"last_updated", existingCoin.LastUpdated,
		).Where(sq.Eq{"coin_name": existingCoin.CoinName}).RunWith(r.db).Exec()
		if err != nil {
			return err
		}
		fmt.Printf("Coin %s updated successfully.", existingCoin.CoinName)
	}
	return err
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

func (r postgresRepo) findCoin(coinName string) (*repo.CoinMarketEntity, error) {
	var existingCoin repo.CoinMarketEntity
	query := sq.Select(
		"coin_name", "price_usd", "last_updated",
	).From(
		"currency",
	).Where(
		sq.Eq{"coin_name": coinName},
	)
	stmt, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row, err := r.db.Query(stmt, args...)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	if err := row.Scan(&existingCoin.CoinName, &existingCoin.PriceUsd, &existingCoin.LastUpdated); err != nil {
		return nil, err
	}
	return &existingCoin, nil
}
