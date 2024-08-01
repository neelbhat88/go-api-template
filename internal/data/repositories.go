package data

// Define repository interfaces here
// e.g. UsersRepository interface which includes all the methods to interact with the users table

//go:generate mockgen -source=repositories.go -destination=mocks/repositories.go -package=mocks

type UsersRepository interface {
	GetUsers() ([]User, error)
}
