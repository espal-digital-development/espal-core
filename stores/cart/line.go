package cart

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: PreSave only allowed Variant OR Bundled
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	cartID             string
	productVariantID   *string
	bundledProductID   *string
	quantity           int
}

// TableName returns the table name that belongs to the current model.
func (l *Line) TableName() string {
	return "CartLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (l *Line) TableAlias() string {
	return "cl"
}
