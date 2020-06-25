package credit

import (
	"time"
)

// Credit database object.
// @synthesize
type Credit struct {
	// TODO :: Implement and bind for migration/foreign keys/indexes
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
func (c *Credit) TableAlias() string {
	return "cre"
}
