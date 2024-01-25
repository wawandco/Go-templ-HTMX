package middleware

import (
	"templ/context"

	"github.com/labstack/echo/v4"
)

func SetCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(&context.Custom{Context: c})
	}
}
