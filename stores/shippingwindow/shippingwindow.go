package shippingwindow

import (
	"time"
)

// ShippingWindow database object.
// @synthesize
type ShippingWindow struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	userGroupID        *string
	startDate          *time.Time
	endDate            *time.Time
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *ShippingWindow) TableAlias() string {
	return "sw"
}
