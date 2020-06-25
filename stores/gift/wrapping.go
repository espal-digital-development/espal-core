package gift

import (
	"time"
)

// Wrapping database object.
// @synthesize
type Wrapping struct {
	// TODO :: Implement (needs price, look, etc.)
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
}

// TableName returns the table name that belongs to the current model.
func (w *Wrapping) TableName() string {
	return "GiftWrapping"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (w *Wrapping) TableAlias() string {
	return "gw"
}
