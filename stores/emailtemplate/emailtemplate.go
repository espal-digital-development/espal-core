package emailtemplate

import (
	"time"
)

// EmailTemplate database object.
// @synthesize
type EmailTemplate struct {
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
}
