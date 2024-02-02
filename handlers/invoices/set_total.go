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

	component := invoicesC.QuantityComponent(invoice.GetTotal(), "Total")
	return render.Render(c, component)
}