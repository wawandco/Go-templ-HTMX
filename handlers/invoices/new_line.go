package invoices

import (
	invoicesC "templ/components/invoices"
	"templ/context"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) NewLine(c echo.Context) error {
	customContext := c.(*context.Custom)
	lines := c.QueryParam("lines")
	invoiceLine := models.InvoiceLine{
		Quantity:  0,
		UnitPrice: 0,
	}
	component := invoicesC.Line(lines, invoiceLine)
	return render.Render(customContext, component)
}
