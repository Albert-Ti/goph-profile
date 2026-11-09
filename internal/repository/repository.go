package repository

type Repository interface {
}

func NewRepository(connString string) (Repository, error) {
	return NewPostgres(connString)
}
