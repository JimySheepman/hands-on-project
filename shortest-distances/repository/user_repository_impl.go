package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type postgreRepository struct {
	postgre *sql.DB
}

func NewPostgreRepository(postgreConnection *sql.DB) UserRepository {
	return &postgreRepository{
		postgre: postgreConnection,
	}
}

func (pr *postgreRepository) Login(ctx context.Context) error {
	return nil
}

func (pr *postgreRepository) Register(ctx context.Context) error {
	return nil
}
