package poll

import (
	"time"
)

// Vote database object.
// @synthesize
type Vote struct {
	// TODO :: Track extra data like IP address?
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	pollOptionID       string
}

// TableName returns the table name that belongs to the current model.
func (vote *Vote) TableName() string {
	return "PollVote"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (vote *Vote) TableAlias() string {
	return "pov"
}
