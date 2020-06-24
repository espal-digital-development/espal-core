package forum

import (
	"time"
)

// PostRating database object.
// @synthesize
type PostRating struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	postID             *string
	score              float64
}

// TableName returns the table name that belongs to the current model.
func (postRating *PostRating) TableName() string {
	return "ForumPostRating"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (postRating *PostRating) TableAlias() string {
	return "fopor"
}
