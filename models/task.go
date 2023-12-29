package models

import "github.com/google/uuid"

type Task struct {
	ID     uuid.UUID `db:"id"`
	Title  string    `db:"title"`
	Active bool      `db:"active"`

	//Timestamps
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type Tasks []Task
