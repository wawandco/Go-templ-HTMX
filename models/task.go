package models

import "github.com/google/uuid"

type Task struct {
	ID     uuid.UUID `db:"id"`
	Title  string    `db:"title"`
	Active bool      `db:"active"`
}

type Tasks []Task
