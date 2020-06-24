package reseller

import (
	"time"
)

// Reseller database object.
// @synthesize
type Reseller struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	active             bool
	country            *uint16
	address            *string // TODO :: Maybe some more detailed fields instead?

	// Images  []Media // TODO :: Implement
}
