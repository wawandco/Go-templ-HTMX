package invoices

import (
	"fmt"
	"strconv"
	"templ/components/invoices"
	"templ/form"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateLines(c echo.Context) error {
	var invoice models.Invoice

	formParams, _ := c.FormParams()
	lineDeleted, _ := strconv.Atoi(formParams.Get("lineDeleted"))

	err := form.Decoder.Decode(&invoice, formParams)
	if err != nil {
		fmt.Println("Error decoding form params: ", err)
		return err
	}

	invoice.InvoiceLines = append(invoice.InvoiceLines[:lineDeleted], invoice.InvoiceLines[lineDeleted+1:]...)

	component := invoices.Lines(invoice.InvoiceLines)
	return render.Render(c, component)
}
