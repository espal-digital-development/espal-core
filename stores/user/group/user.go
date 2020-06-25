package group

import (
	"time"
)

// User database object.
// @synthesize
type User struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	userGroupID        string
	userID             string
}

// TableName returns the table name that belongs to the current model.
func (u *User) TableName() string {
	return "UserGroupsUser"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (u *User) TableAlias() string {
	return "ugu"
}
