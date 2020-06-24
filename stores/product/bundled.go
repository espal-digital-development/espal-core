package product

import (
	"time"
)

// Bundled database object.
// @synthesize
type Bundled struct {
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
	variantsCanBeSoldSeperately bool
	taxGroupID                  string
	nameRepresentationID        *string // Property
	descriptionRepresentationID *string // Property
	imageRepresentationID       *string // Property

	// BundledVariants       []BundledVariant
	// BundledReviews []BundledReview
}

// TableName returns the table name that belongs to the current model.
func (bundled *Bundled) TableName() string {
	return "BundledProduct"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (bundled *Bundled) TableAlias() string {
	return "bpr"
}
