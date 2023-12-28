package tasks

import (
	tasksC "templ/components/tasks"
	"templ/internal/tasks"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Count(c echo.Context) error {
	count, err := tasks.Count(h.DB)
	if err != nil {
		return err
	}

	return render.Render(c, tasksC.Count(count))
}
