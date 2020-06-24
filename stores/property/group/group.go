package group

import (
	"time"
)

// Group database object.
// @synthesize
type Group struct {
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

	// Properties []Property
}

// TableName returns the table name that belongs to the current model.
func (group *Group) TableName() string {
	return "PropertyGroup"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (group *Group) TableAlias() string {
	return "prgr"
}
