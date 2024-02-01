package invoices

import (
	invoicesC "templ/components/invoices"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) New(c echo.Context) error {
	invoice := models.Invoice{
		InvoiceLines: []models.InvoiceLine{
			{
				Quantity:  0,
				UnitPrice: 0,
			},
		},
	}
	return render.Render(c, invoicesC.New(invoice))
}
