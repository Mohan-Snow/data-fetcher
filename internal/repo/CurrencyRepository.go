package repo

import (
	"data-fetcher/internal"
	"data-fetcher/internal/model/repo"
	"database/sql"
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

func (r postgresRepo) Save(entity *repo.CmEntity) error {
	toSql, args, err := sq.Insert("currency").
		Columns("coin_name", "price_usd", "last_updated").
		Values(entity.CoinName, entity.PriceUsd, entity.LastUpdated).
		RunWith(r.db).
		ToSql()
	if err != nil {
		return err
	}

	fmt.Println(toSql)
	fmt.Println(args)

	//_, err = r.db.Query(toSql, args...)
	//_, err = r.db.Exec(toSql, args...)
	//conn, err := r.db.Conn(context.Background())
	//_, err = conn.QueryContext(context.Background(), toSql, args...)

	//_, err := r.db.Exec("INSERT INTO currency(COIN_NAME, PRICE_USD, LAST_UPDATED) VALUES ($1,$2,$3)",
	//	entity.CoinName, entity.PriceUsd, entity.LastUpdated)
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
