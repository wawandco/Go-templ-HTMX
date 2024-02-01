package invoices

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"templ/form"
	"templ/models"
	"text/template"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/labstack/echo/v4"
)

//go:embed test.html
var test string

func (h *Handler) Download(c echo.Context) error {

	var invoice models.Invoice
	formParams, _ := c.FormParams()
	err := form.Decoder.Decode(&invoice, formParams)
	if err != nil {
		fmt.Println("Error decoding form params: ", err)
		return err
	}

	template, err := template.New("invoice").Funcs(template.FuncMap{
		"getTotal": invoice.GetTotal,
		"lineTotal": func(line models.InvoiceLine) int {
			return line.Quantity * line.UnitPrice
		},
		"subTotal": invoice.GetSubtotal,
	}).Parse(test)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return err
	}

	var buffer1 bytes.Buffer
	err = template.Execute(&buffer1, invoice)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		return err
	}

	page := rod.New().MustConnect().MustPage()
	page.SetDocumentContent(buffer1.String())

	page.MustWaitLoad()

	bt, err := page.PDF(&proto.PagePrintToPDF{})
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, bt)
	if err != nil {
		return err
	}

	defer bt.Close()

	c.Response().Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")
	c.Response().Header().Set("HX-Redirect", "invoice.pdf")
	c.Response().Write(buffer.Bytes())

	return nil
}
