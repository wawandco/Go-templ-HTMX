package tasks

import (
	"database/sql"
	"templ/models"
)

func List(db *sql.DB) (models.Tasks, error) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks models.Tasks
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Active); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}
