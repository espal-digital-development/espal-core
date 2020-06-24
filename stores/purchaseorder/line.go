package purchaseorder

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: How to know the products that are being bought? Can link Variant/Bundled, but not unknown
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	sorting            uint
	purchaseOrderID    string
	quantity           int
	price              float32 // Price per unit
	vatPercentage      float32
	comments           *string
}

// TableName returns the table name that belongs to the current model.
func (line *Line) TableName() string {
	return "PurchaseOrderLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (line *Line) TableAlias() string {
	return "porl"
}
