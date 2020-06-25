package task

import (
	"time"
)

// Task database object.
// @synthesize
type Task struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	issuedByID         string  // User
	assignedToID       *string // User
	description        string
	completedNotes     *string
	completedAt        *time.Time
}

// TableName returns the table name that belongs to the current model.
func (t *Task) TableName() string {
	return "UserTask"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (t *Task) TableAlias() string {
	return "ut"
}
