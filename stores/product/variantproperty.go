package product

import (
	"time"
)

// VariantProperty database object.
// @synthesize
type VariantProperty struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	variantID          string
	propertyID         string
	sorting            uint
	key                *string
}

// TableName returns the table name that belongs to the current model.
func (variantProperty *VariantProperty) TableName() string {
	return "ProductVariantProperty"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (variantProperty *VariantProperty) TableAlias() string {
	return "provp"
}
