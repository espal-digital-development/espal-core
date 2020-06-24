package cart

import (
	"time"
)

// Cart database object.
// @synthesize
type Cart struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	domainID           string
	userID             *string

	// Lines []CartLine
}

// TableAlias returns the unique resolved table alias for use in queries.
func (cart *Cart) TableAlias() string {
	return "car"
}
