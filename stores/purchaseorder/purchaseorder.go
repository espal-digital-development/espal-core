package purchaseorder

import (
	"time"
)

// PurchaseOrder database object.
// @synthesize
type PurchaseOrder struct {
	// TODO :: Enable and link Supplier
	// TODO :: Link more info here?
	id                    string
	createdByID           string
	updatedByID           *string
	createdAt             time.Time
	updatedAt             *time.Time
	createdByFirstName    *string
	createdBySurname      *string
	updatedByFirstName    *string
	updatedBySurname      *string
	supplierID            string
	currency              uint16
	comments              *string
	sellingPartyAutograph *string // bytes for autograph image
	buyingPartyAutograph  *string // bytes for autograph image

	// Lines []Line
}

// TableAlias returns the unique resolved table alias for use in queries.
func (p *PurchaseOrder) TableAlias() string {
	return "por"
}
