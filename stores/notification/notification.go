package notification

import (
	"time"
)

const (
	TargetDomain = "1"
	TargetSite   = "2"
	TargetUser   = "3"
	TargetSlug   = "4"
)

// Notification database object.
// @synthesize
type Notification struct {
	id                 string
	createdByID        *string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	target             string
	key                string
	value              string
}
