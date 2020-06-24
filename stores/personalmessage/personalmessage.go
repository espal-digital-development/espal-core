package personalmessage

import (
	"time"
)

// PersonalMessage database object.
// @synthesize
type PersonalMessage struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	userID             string
	recipientID        string
	responseToID       *string // PersonalMessage
	title              *string
	message            string
}

// TableAlias returns the unique resolved table alias for use in queries.
func (personalMessage *PersonalMessage) TableAlias() string {
	return "pme"
}
