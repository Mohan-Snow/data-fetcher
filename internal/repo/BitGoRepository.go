package repo

import (
	"data-fetcher/internal"
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

func (r postgresRepo) Save(data string) error {
	_, err := r.db.Exec("INSERT INTO test_data(data) VALUES ($1)", data)
	if err != nil {
		return err
	}
	return nil
}

func (r postgresRepo) GetById(id int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r postgresRepo) GetAll() (string, error) {
	//TODO implement me
	panic("implement me")
}
