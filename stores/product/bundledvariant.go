package product

import (
	"time"
)

// BundledVariant database object.
// @synthesize
type BundledVariant struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	bundledID          string
	variantID          string
	quantity           uint
}

// TableName returns the table name that belongs to the current model.
func (b *BundledVariant) TableName() string {
	return "BundledProductVariant"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (b *BundledVariant) TableAlias() string {
	return "bprv"
}
