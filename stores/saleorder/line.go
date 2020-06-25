package saleorder

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: PreSave/PreUpdate only allowed Variant OR Bundled OR none (custom line)
	// TODO :: Link this to ShipmentLine(s)
	// TODO :: Discount amounts (percentual/fixed) should be separate lines here?
	// TODO :: ShipmentLine?
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
	saleOrderID        string
	deliveryMethodID   string
	productVariantID   *string
	bundledProductID   *string
	quantity           int
	price              float32 // Price per unit
	vatPercentage      float32
	comments           *string
}

// TableName returns the table name that belongs to the current model.
func (l *Line) TableName() string {
	return "SaleOrderLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (l *Line) TableAlias() string {
	return "sol"
}
