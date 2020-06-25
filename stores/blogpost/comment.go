package blogpost

import (
	"time"
)

// Comment database object.
// @synthesize
type Comment struct {
	// TODO :: Title is optional, but should need a setting.
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	blogPostID         string
	title              *string
	message            string
}

// TableName returns the table name that belongs to the current model.
func (c *Comment) TableName() string {
	return "BlogPostComment"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (c *Comment) TableAlias() string {
	return "bpc"
}
