package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=templ sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
