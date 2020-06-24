package product

import (
	"time"
)

// Model database object.
// @synthesize
type Model struct {
	id                          string
	createdByID                 string
	updatedByID                 *string
	createdAt                   time.Time
	updatedAt                   *time.Time
	createdByFirstName          *string
	createdBySurname            *string
	updatedByFirstName          *string
	updatedBySurname            *string
	active                      bool
	sorting                     uint
	key                         *string
	taxGroupID                  string
	nameRepresentationID        *string // Property
	descriptionRepresentationID *string // Property
	imageRepresentationID       *string // Property

	// Variants     []Variant
	// ModelReviews []ModelReview
}

// TableName returns the table name that belongs to the current model.
func (model *Model) TableName() string {
	return "ProductModel"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (model *Model) TableAlias() string {
	return "prom"
}
