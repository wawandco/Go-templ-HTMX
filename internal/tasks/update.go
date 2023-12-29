package tasks

import "database/sql"

func Update(db *sql.DB, t Task) error {
	_, err := db.Exec("UPDATE tasks SET title = $1, active = $2 WHERE id = $3", t.Title, t.Active, t.ID)
	if err != nil {
		return err
	}

	return nil
}
