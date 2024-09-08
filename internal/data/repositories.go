package data

// Define repository interfaces here
// e.g. UsersRepository interface which includes all the methods to interact with the users table

type UsersRepository interface {
	GetUsers() ([]User, error)
}
