package v1

import (
	"github.com/jmoiron/sqlx"
	"github.com/neelbhat88/go-api-template/internal/data/postgres"
)

type Handler struct {
	DB postgres.DB
}

func NewHandler(db *sqlx.DB) Handler {
	database := postgres.NewPostgresDB(db)

	return Handler{
		DB: database,
	}
}
