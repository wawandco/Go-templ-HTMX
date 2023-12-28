package tasks

import (
	"database/sql"
	"templ/models"

	"github.com/google/uuid"
)

func Find(db *sql.DB, ID uuid.UUID) (models.Task, error) {
	var task models.Task
	err := db.QueryRow("SELECT * FROM tasks WHERE id = $1", ID).Scan(&task.ID, &task.Title, &task.Active)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
