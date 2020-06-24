package media

import (
	"time"
)

// Media database object.
// @synthesize
type Media struct {
	// TODO :: More fields?
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
	filePath           string
}
