package stocklocation

import (
	"time"
)

// Mutation database object.
// @synthesize
type Mutation struct {
	// TODO :: Keep track if it's a manual change or by
	// SaleOrder/Shipment/PurchaseOrder/Receiving etc. and their ID.
	id                    string
	createdByID           string
	updatedByID           *string
	createdAt             time.Time
	updatedAt             *time.Time
	createdByFirstName    *string
	createdBySurname      *string
	updatedByFirstName    *string
	updatedBySurname      *string
	sourceID              *string // StockLocation: Filled it came from another location.
	targetID              string  // StockLocation
	productVariantID      string
	modifier              int // e.g. (+)10 or -15
	ballanceAfterModifier int
	comments              *string
}

// TableName returns the table name that belongs to the current model.
func (m *Mutation) TableName() string {
	return "StockMutation"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (m *Mutation) TableAlias() string {
	return "slm"
}
