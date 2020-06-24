package shipment

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: PreSave/PreUpdate only allowed Variant OR Bundled OR none (custom line)
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
	shipmentID         string
	saleOrderLineID    *string
	productVariantID   *string
	bundledProductID   *string
	quantity           int
	comments           *string
}

// TableName returns the table name that belongs to the current model.
func (line *Line) TableName() string {
	return "ShipmentLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (line *Line) TableAlias() string {
	return "shl"
}
