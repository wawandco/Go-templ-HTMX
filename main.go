package main

import (
	"templ/actions/users"
	"templ/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db := db.GetConnection()

	app := echo.New()

	userHandler := users.Handler{DB: db}
	app.GET("/", userHandler.List)
	app.Start(":8080")
}
