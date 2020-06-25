package poll

import (
	"time"
)

// Option database object.
// @synthesize
type Option struct {
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
	sorting            uint
	pollID             string

	// Votes []Vote
}

// TableName returns the table name that belongs to the current model.
func (o *Option) TableName() string {
	return "PollOption"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (o *Option) TableAlias() string {
	return "poo"
}
