package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID           uuid.UUID     `form:"ID"`
	Reference    string        `form:"Reference"`
	Date         time.Time     `orm:"Date"`
	Total        float64       `form:"Total"`
	Discount     float64       `form:"Discount"`
	Subtotal     float64       `form:"Subtotal"`
	InvoiceLines []InvoiceLine `form:"InvoiceLines"`
}

type Invoices []Invoice

func (i *Invoice) GetTotal() string {
	i.Subtotal = i.GetSubtotal()

	if i.Discount > 0 {
		i.Subtotal = i.Subtotal - (i.Subtotal * (i.Discount / 100))
	}

	i.Total = i.Subtotal
	return fmt.Sprintf("$%.2f", i.Total)
}

func (i *Invoice) GetSubtotal() float64 {
	var total int
	for _, line := range i.InvoiceLines {
		total = total + (line.Quantity * line.UnitPrice)
	}

	i.Subtotal = float64(total)
	return i.Subtotal
}
