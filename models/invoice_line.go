package models

import (
	"fmt"

	"github.com/google/uuid"
)

type InvoiceLine struct {
	ID          uuid.UUID `form:"ID"`
	InvoiceID   uuid.UUID `form:"InvoiceID"`
	Description string    `form:"Description"`
	Quantity    int       `form:"Quantity"`
	UnitPrice   int       `form:"UnitPrice"`
}

func (i *InvoiceLine) Total() string {
	total := i.Quantity * i.UnitPrice

	return fmt.Sprintf("$%d", total)
}

func (i *InvoiceLine) UnitPriceString() string {
	return fmt.Sprintf("%d", i.UnitPrice)
}

type InvoiceLines []InvoiceLine
