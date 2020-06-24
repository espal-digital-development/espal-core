package newsletter

import (
	"time"
)

// Newsletter database object.
// @synthesize
type Newsletter struct {
	// TODO :: Make time-interval checker if SendAtTime is set to send in the future.
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
	active             bool
	sendAt             *time.Time
}
