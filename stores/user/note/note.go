package note

import (
	"time"
)

// Note database object.
// @synthesize
type Note struct {
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
	title              *string
	contents           string
}

// TableName returns the table name that belongs to the current model.
func (note *Note) TableName() string {
	return "UserNote"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (note *Note) TableAlias() string {
	return "un"
}
