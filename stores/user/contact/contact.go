package contact

import (
	"time"
)

// Contact database object.
// @synthesize
type Contact struct {
	// TODO :: Make a setting to mask this as `Friend`?
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	userID             *string
	contactID          string
	sorting            uint
	comments           *string // Notes on the ContactUser
	contactFirstName   *string
	contactSurname     *string
}

// TableName returns the table name that belongs to the current model.
func (contact *Contact) TableName() string {
	return "UserContact"
}
