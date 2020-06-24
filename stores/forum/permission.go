package forum

import (
	"time"
)

// Permissions are used to identify actions Users can
// perform on Forum instances.
const (
	PermissionPost uint8 = iota + 1
	PermissionReply
	PermissionEdit
	PermissionDelete
)

// Permission database object.
// @synthesize
type Permission struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	forumID            *string
	userID             *string
	permission         uint8
}

// TableName returns the table name that belongs to the current model.
func (permission *Permission) TableName() string {
	return "ForumPermission"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (permission *Permission) TableAlias() string {
	return "fop"
}
