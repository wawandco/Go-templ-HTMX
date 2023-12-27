package main

import (
	"templ/db"
	"templ/handlers/tasks"

	"github.com/labstack/echo/v4"
)

func main() {
	db := db.GetConnection()

	app := echo.New()

	taskHandler := tasks.Handler{DB: db}
	app.GET("/", taskHandler.List)
	app.Start(":8080")
}
