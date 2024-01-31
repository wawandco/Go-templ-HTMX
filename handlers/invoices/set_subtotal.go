package invoices

import (
	"fmt"
	invoicesC "templ/components/invoices"
	"templ/form"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetSubtotal(c echo.Context) error {
	var invoice models.Invoice

	formParams, _ := c.FormParams()
	err := form.Decoder.Decode(&invoice, formParams)
	if err != nil {
		return err
	}

	var subTotal float64
	for _, line := range invoice.InvoiceLines {
		subTotal += float64(line.Quantity) * line.UnitPrice
	}

	component := invoicesC.QuantityComponent(fmt.Sprintf("$%.2f", subTotal), "Subtotal")
	return render.Render(c, component)
}
