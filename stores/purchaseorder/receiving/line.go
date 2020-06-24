package receiving

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: When filling ProductVariant or BundledProduct; only one is allowed.
	// PurchaseOrderLine can be null when something wasn't linked to one.
	// If none of the above is true; products can still be linked individually.
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	receivingID        string
	lineID             *string
	productVariantID   *string
	bundledProductID   *string
	quantity           uint
	comments           *string
}

// TableName returns the table name that belongs to the current model.
func (line *Line) TableName() string {
	return "PurchaseOrderReceivingLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (line *Line) TableAlias() string {
	return "porrl"
}
