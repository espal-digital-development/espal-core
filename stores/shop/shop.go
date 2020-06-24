package shop

import (
	"time"
)

// Shop database object.
// @synthesize
type Shop struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	language           *uint16
	currencies         string
}

// TableAlias returns the unique resolved table alias for use in queries.
func (shop *Shop) TableAlias() string {
	return "sho"
}
