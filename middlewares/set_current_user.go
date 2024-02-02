package middleware

import (
	"context"
	"templ/models"

	"github.com/labstack/echo/v4"
)

func SetCurrentUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{
			FirstName: "Messi",
			LastName:  "Caicedo",
		}

		context := context.WithValue(c.Request().Context(), "user", user)
		c.SetRequest(c.Request().WithContext(context))

		return next(c)
	}
}
