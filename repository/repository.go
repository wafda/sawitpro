// This file contains the repository implementation layer.
package repository

import (
	"database/sql"

	"github.com/go-playground/validator/v10"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db       *sql.DB
	validate *validator.Validate
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}

	return &Repository{
		Db:       db,
		validate: validator.New(),
	}
}
