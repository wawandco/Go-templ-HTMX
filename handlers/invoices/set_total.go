package invoices

import (
	"fmt"
	invoicesC "templ/components/invoices"
	"templ/form"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetTotal(c echo.Context) error {
	var invoice models.Invoice

	formParams, _ := c.FormParams()

	err := form.Decoder.Decode(&invoice, formParams)
	if err != nil {
		fmt.Println("Error decoding form params: ", err)
		return err
	}

	//Sum the total of the invoice lines
	var total float64
	for _, line := range invoice.InvoiceLines {
		total += float64(line.Quantity) * line.UnitPrice
	}

	descount := invoice.Discount
	if descount > 0 {
		total = total - (total * (descount / 100))
	}

	component := invoicesC.QuantityComponent(fmt.Sprintf("$%.2f", total), "Total")
	return render.Render(c, component)
}
