package tasks

import (
	"database/sql"

	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID
	Title  string
	Active bool
}

func Create(db *sql.DB, t Task) error {
	_, err := db.Exec("INSERT INTO tasks (id, title, active) VALUES ($1, $2, $3)", t.ID, t.Title, t.Active)
	if err != nil {
		return err
	}

	return nil
}
