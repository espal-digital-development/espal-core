package price

import (
	"time"
)

// Mutation database object.
// @synthesize
type Mutation struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	productVariantID   string
	price              float64
}

// TableName returns the table name that belongs to the current model.
func (m *Mutation) TableName() string {
	return "PriceMutation"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (m *Mutation) TableAlias() string {
	return "prm"
}
