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
	app.POST("/tasks", taskHandler.Create)
	app.GET("/tasks_count", taskHandler.Count)
	app.DELETE("/tasks/:id", taskHandler.Delete)
	app.GET("/tasks/:id", taskHandler.Get)
	app.Start(":8080")
}
