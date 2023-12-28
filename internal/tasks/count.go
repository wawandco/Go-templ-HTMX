package tasks

import (
	"database/sql"
)

func Count(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
