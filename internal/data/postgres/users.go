package postgres

import (
	"github.com/neelbhat88/go-api-template/internal/data"
	"github.com/rs/zerolog/log"
)

func (db DB) GetUsers() ([]data.User, error) {
	var users []data.User
	err := db.Select(&users, `
		SELECT id, email
		FROM users
	`)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get users")
		return nil, err
	}

	return users, nil
}
