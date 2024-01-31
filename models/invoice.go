package models

import (
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
