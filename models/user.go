package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"name"`
	LastName  string    `db:"last_name"`
}

type Users []User
