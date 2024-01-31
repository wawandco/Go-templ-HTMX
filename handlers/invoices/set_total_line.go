package invoices

import (
	"fmt"
	"strconv"
	invoicesC "templ/components/invoices"
	"templ/models"
	"templ/render"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetTotalLine(c echo.Context) error {
	var line models.InvoiceLine
	lineParam := c.QueryParam("line")

	formParams, _ := c.FormParams()

	line.Quantity, _ = strconv.Atoi(formParams.Get("InvoiceLines[" + lineParam + "].Quantity"))
	line.UnitPrice, _ = strconv.ParseFloat(formParams.Get("InvoiceLines["+lineParam+"].UnitPrice"), 64)

	total := float64(line.Quantity) * line.UnitPrice

	component := invoicesC.TotalLine(fmt.Sprintf("$%.2f", total))
	return render.Render(c, component)
}
