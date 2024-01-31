package internal

type PostgresRepo interface {
	Save(data string) error
	GetById(id int) (string, error)
	GetAll() (string, error)
}

type IFetchDataService interface {
	FetchDataAndSave(symbol string, token string) (string, error)
}
