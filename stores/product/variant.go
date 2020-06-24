package product

import (
	"time"
)

// Variant database object.
// @synthesize
type Variant struct {
	id                          string
	createdByID                 string
	updatedByID                 *string
	createdAt                   time.Time
	updatedAt                   *time.Time
	createdByFirstName          *string
	createdBySurname            *string
	updatedByFirstName          *string
	updatedBySurname            *string
	modelID                     string
	active                      bool
	key                         *string
	sorting                     uint
	taxGroupID                  string
	nameRepresentationID        *string // Property
	descriptionRepresentationID *string // Property
	imageRepresentationID       *string // Property

	// BundledVariants []BundledVariant
	// Reviews         []VariantReview
}

// TableName returns the table name that belongs to the current model.
func (variant *Variant) TableName() string {
	return "ProductVariant"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (variant *Variant) TableAlias() string {
	return "prov"
}
