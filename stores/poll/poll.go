package poll

import (
	"time"
)

// Poll database object.
// @synthesize
type Poll struct {
	id                   string
	createdByID          string
	updatedByID          *string
	createdAt            time.Time
	updatedAt            *time.Time
	createdByFirstName   *string
	createdBySurname     *string
	updatedByFirstName   *string
	updatedBySurname     *string
	active               bool
	startDate            *time.Time
	endDate              *time.Time
	allowAnonymousVoting bool
	comments             *string

	// Options []PollOption
}
