package main

import (
	"templ/db"
	"templ/handlers/tasks"
	"templ/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db := db.GetConnection()
	app := echo.New()
	app.Use(middleware.SetCustomContext)

	taskHandler := tasks.Handler{DB: db}
	app.GET("/", taskHandler.List)
	app.POST("/tasks", taskHandler.Create)
	app.GET("/tasks_count", taskHandler.Count)
	app.DELETE("/tasks/:id", taskHandler.Delete)
	app.GET("/tasks/:id", taskHandler.Get)
	app.PUT("/tasks/:id", taskHandler.Update)
	app.GET("/active_tasks", taskHandler.ActiveList)
	app.GET("/completed_tasks", taskHandler.CompletedList)
	app.PATCH("/tasks/:id/update_status", taskHandler.UpdateStatus)
	app.Start(":8080")
}
