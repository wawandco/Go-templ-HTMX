package invoices

import (
	invoicesC "templ/components/invoices"
	"templ/context"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) NewLine(c echo.Context) error {
	customContext := c.(*context.Custom)
	lines := c.QueryParam("lines")
	component := invoicesC.Line(lines)
	return render.Render(customContext, component)
}
