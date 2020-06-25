package blogpost

import (
	"time"
)

// Like database object.
// @synthesize
type Like struct {
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
}

// TableName returns the table name that belongs to the current model.
func (l *Like) TableName() string {
	return "BlogPostLike"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (l *Like) TableAlias() string {
	return "bpl"
}
