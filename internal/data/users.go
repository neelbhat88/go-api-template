package data

type User struct {
	ID    int    `db:"id"`
	Email string `db:"email"`
}
