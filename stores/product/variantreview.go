package product

import (
	"time"
)

// VariantReview database object.
// @synthesize
type VariantReview struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	// TODO :: Setting if review by User/-group is required or auto-accept.
	reviewedByID   *string // User
	reviewedOnDate *time.Time
	reviewNotes    *string
	approved       *bool
	// More flexible; e.g. 80% can represent 4 out of 5 stars.
	rating      float32
	title       string
	description string
	variantID   string
}

// TableName returns the table name that belongs to the current model.
func (v *VariantReview) TableName() string {
	return "ProductVariantReview"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (v *VariantReview) TableAlias() string {
	return "provr"
}
