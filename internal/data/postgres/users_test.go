package postgres

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB_GetUsers(t *testing.T) {
	db, cleanup, err := CreateTestDatabase()
	defer cleanup()
	pg := NewPostgresDB(db)

	_, err = pg.DB.Exec(`
		INSERT INTO users (id, email, password) VALUES (1, 'neel@test.com', 'password');
	`)
	if err != nil {
		t.Fatal("Failed to insert user")
	}

	users, err := pg.GetUsers()
	if err != nil {
		t.Fatalf("Failed to get users: %v", err)
	}

	assert.Len(t, users, 1)
}

func TestDB_GetUsersWithNoUsersInDB(t *testing.T) {
	db, cleanup, err := CreateTestDatabase()
	defer cleanup()
	pg := NewPostgresDB(db)

	users, err := pg.GetUsers()
	if err != nil {
		t.Fatalf("Failed to get users: %v", err)
	}

	assert.Len(t, users, 0)
}
