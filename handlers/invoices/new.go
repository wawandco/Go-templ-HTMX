package invoices

import (
	invoicesC "templ/components/invoices"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) New(c echo.Context) error {
	return render.Render(c, invoicesC.New())
}
