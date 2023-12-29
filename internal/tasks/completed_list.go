package tasks

import (
	"database/sql"
	"templ/models"
)

func CompletedList(db *sql.DB) (models.Tasks, error) {
	rows, err := db.Query("SELECT * FROM tasks WHERE active = false ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	var tasks models.Tasks
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Active, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}
