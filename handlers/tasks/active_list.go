package tasks

import (
	tasksC "templ/components/tasks"
	"templ/internal/tasks"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ActiveList(c echo.Context) error {
	tsks, err := tasks.ActiveList(h.DB)
	if err != nil {
		return err
	}

	if c.Request().Header.Get("Hx-Request") == "true" {
		return render.Render(c, tasksC.Table(tsks, "/active_tasks"))
	}

	return render.Render(c, tasksC.List(tsks, "/active_tasks"))
}
