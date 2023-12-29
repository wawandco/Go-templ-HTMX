package tasks

import (
	"net/http"
	"templ/internal/tasks"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateStatus(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	tsk, err := tasks.Find(h.DB, id)
	if err != nil {
		return err
	}

	_, err = tasks.UpdateStatus(h.DB, tsk.ID, !tsk.Active)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
