package usersAdmin

import (
	"github.com/neelbhat88/go-api-template/internal/data"
	"github.com/neelbhat88/go-api-template/internal/entities"
)

func LoadAllUsers(userRepo data.UsersRepository) ([]entities.User, error) {
	dbUsers, err := userRepo.GetUsers()
	if err != nil {
		return []entities.User{}, err
	}

	var users []entities.User
	for _, dbUser := range dbUsers {
		users = append(users, entities.User{
			ID:    dbUser.ID,
			Email: dbUser.Email,
		})
	}

	return users, nil
}
