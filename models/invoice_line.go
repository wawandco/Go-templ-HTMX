package models

import "github.com/google/uuid"

type InvoiceLine struct {
	ID          uuid.UUID `form:"ID"`
	InvoiceID   uuid.UUID `form:"InvoiceID"`
	Description string    `form:"Description"`
	Quantity    int       `form:"Quantity"`
	UnitPrice   float64   `form:"UnitPrice"`
}

type InvoiceLines []InvoiceLine
