package main

import (
	"templ/handlers/invoices"

	"templ/middleware"

	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	app.Use(middleware.SetCustomContext)

	invoicesHandlers := invoices.Handler{}
	app.GET("/", invoicesHandlers.New)
	app.GET("/new-line", invoicesHandlers.NewLine)
	app.DELETE("/remove-line", invoicesHandlers.RemoveLine)
	app.POST("/set-total-line", invoicesHandlers.SetTotalLine)
	app.POST("/set-subtotal", invoicesHandlers.SetSubtotal)
	app.POST("/set-total", invoicesHandlers.SetTotal)

	app.Start(":8080")
}
