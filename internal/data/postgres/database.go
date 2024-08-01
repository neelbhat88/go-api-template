package postgres

import (
	"embed"
	"github.com/jmoiron/sqlx"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

type DB struct {
	*sqlx.DB
}

func GetMigrations() PostgresMigrations {
	return PostgresMigrations{
		SchemaName:     "public",
		MigrationFiles: migrationFiles,
		Path:           "migrations",
	}
}

func NewPostgresDB(db *sqlx.DB) DB {
	postgresDB := DB{
		DB: db,
	}

	return postgresDB
}
