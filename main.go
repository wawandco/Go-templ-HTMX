package main

import (
	"net/http"
	"templ/handlers/invoices"

	"templ/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	app.Use(middlewares.SetCustomContext)

	invoicesHandlers := invoices.Handler{}
	app.GET("/", invoicesHandlers.New)
	app.GET("/new-line", invoicesHandlers.NewLine)
	app.DELETE("/remove-line", invoicesHandlers.RemoveLine)
	app.POST("/set-total-line", invoicesHandlers.SetTotalLine)
	app.POST("/set-subtotal", invoicesHandlers.SetSubtotal)
	app.POST("/set-total", invoicesHandlers.SetTotal)
	app.PUT("/update-lines", invoicesHandlers.UpdateLines)
	app.GET("/download", invoicesHandlers.Download)

	// serve static files
	fs := http.FileServer(http.Dir("assets"))
	app.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fs)))
	app.Start(":8080")
}
