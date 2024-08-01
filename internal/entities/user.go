package entities

// User represents a user entity
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`

	IsAdmin bool `json:"isAdmin"`
}
