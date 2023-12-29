package tasks

import (
	"strconv"
	tasksC "templ/components/tasks"
	"templ/internal/tasks"
	"templ/models"
	"templ/render"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Update(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))
	index := c.Request().URL.Query().Get("index")
	title := c.FormValue("titleUpdate")
	t, err := tasks.Find(h.DB, id)
	if err != nil {
		return err
	}

	if title == "" {
		return render.Render(c, tasksC.Edit(models.Task{
			ID:     t.ID,
			Title:  "",
			Active: t.Active,
		}, index, "Title is required"))
	}

	t.Title = title
	err = tasks.Update(h.DB, tasks.Task{
		ID:     t.ID,
		Title:  t.Title,
		Active: t.Active,
	})
	if err != nil {
		return err
	}

	idx, _ := strconv.Atoi(index)
	return render.Render(c, tasksC.Item(t, idx))
}
