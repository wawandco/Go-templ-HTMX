package invoices

import (
	invoicesC "templ/components/invoices"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) NewLine(c echo.Context) error {

	// value := c.Request().Context().Value("user").(models.User)
	// fmt.Println(value)

	lines := c.QueryParam("lines")
	invoiceLine := models.InvoiceLine{
		Quantity:  0,
		UnitPrice: 0,
	}
	component := invoicesC.Line(lines, invoiceLine)
	return render.Render(c, component)
}
