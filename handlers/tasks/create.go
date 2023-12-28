package tasks

import (
	tasksC "templ/components/tasks"
	"templ/internal/tasks"
	"templ/render"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	title := c.FormValue("title")

	if title == "" {
		return render.Render(c, tasksC.New("Title is required"))
	}

	if err := tasks.Create(h.DB, tasks.Task{
		ID:     uuid.New(),
		Title:  title,
		Active: true,
	}); err != nil {
		return err
	}

	return render.Render(c, tasksC.New())
}
