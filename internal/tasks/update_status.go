package tasks

import (
	"database/sql"
	"templ/models"

	"github.com/google/uuid"
)

func UpdateStatus(db *sql.DB, id uuid.UUID, status bool) (models.Task, error) {
	var t models.Task
	err := db.QueryRow("UPDATE tasks SET active = $1 WHERE id = $2 RETURNING *", status, id).Scan(&t.ID, &t.Title, &t.Active, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return models.Task{}, err
	}

	return t, nil
}
