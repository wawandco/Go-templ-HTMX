package tasks

import (
	"database/sql"

	"github.com/google/uuid"
)

func Delete(db *sql.DB, id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
