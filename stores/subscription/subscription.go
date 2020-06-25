package subscription

import (
	"time"
)

// Subscription database object.
// @synthesize
type Subscription struct {
	// TODO :: Implement as Product with a pay-model or as free
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

// TableAlias returns the unique resolved table alias for use in queries.
func (s *Subscription) TableAlias() string {
	return "sub"
}
