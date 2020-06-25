package product

import (
	"time"
)

// BundledProperty database object.
// @synthesize
type BundledProperty struct {
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
	key                *string
	bundledProductID   string
	propertyID         string
}

// TableName returns the table name that belongs to the current model.
func (b *BundledProperty) TableName() string {
	return "BundledProductProperty"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (b *BundledProperty) TableAlias() string {
	return "bprp"
}
