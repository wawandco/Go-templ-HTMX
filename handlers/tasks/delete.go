package tasks

import (
	"templ/internal/tasks"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Delete(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))
	err := tasks.Delete(h.DB, id)
	if err != nil {
		return err
	}

	return c.NoContent(200)
}
