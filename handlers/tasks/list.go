package tasks

import (
	users "templ/components/tasks"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) List(c echo.Context) error {
	usrs := models.Tasks{}

	return render.Render(c, users.List(usrs))
}
