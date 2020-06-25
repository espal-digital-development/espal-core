package receiving

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: When filling ProductVariant or BundledProduct; only one is allowed.
	// ReturnOrderLine can be null when something gets returned that wasn't
	// notified upfront. SaleOrderLine can still be optionally linked.
	// If none of the above is true; variants can still be linked individually.
	id                     string
	createdByID            string
	updatedByID            *string
	createdAt              time.Time
	updatedAt              *time.Time
	createdByFirstName     *string
	createdBySurname       *string
	updatedByFirstName     *string
	updatedBySurname       *string
	returnOrderReceivingID string
	returnOrderLineID      *string
	saleOrderLineID        *string
	productVariantID       *string
	bundledProductID       *string
	quantity               uint
	comments               *string
}

// TableName returns the table name that belongs to the current model.
func (l *Line) TableName() string {
	return "ReturnOrderReceivingLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (l *Line) TableAlias() string {
	return "rorl"
}
