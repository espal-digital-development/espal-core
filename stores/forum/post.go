package forum

import (
	"time"
)

// Post database object.
// @synthesize
type Post struct {
	// TODO :: Sticky should only be setable by someone with moderating rights.
	// TODO :: Delete/Active (hiding/deleting a post)
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	forumID            string
	responseToID       *string // ForumPost
	sticky             bool
	title              *string
	message            string
	timesEdited        *uint
	name               string

	// Ratings []PostRating
}

// TableName returns the table name that belongs to the current model.
func (post *Post) TableName() string {
	return "ForumPost"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (post *Post) TableAlias() string {
	return "fopo"
}
