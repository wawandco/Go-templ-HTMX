package users

import (
	"templ/layouts/users"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) List(c echo.Context) error {
	usrs := []models.User{
		{Name: "John"},
		{Name: "Jane"},
		{Name: "Ronaldo"},
	}

	return render.Render(c, users.List(usrs))
}
