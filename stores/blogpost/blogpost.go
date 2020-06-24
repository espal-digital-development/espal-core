package blogpost

import (
	"time"
)

// BlogPost database object.
// @synthesize
type BlogPost struct {
	// TODO :: Approval is optional, but should need a setting.
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
	sectionID          *string
	// approvedBy     user.UserEntity
	approvedByID   *string
	approvedDate   *time.Time
	publishDate    *time.Time
	expirationDate *time.Time
	comments       *string

	// Likes []BlogPostLike
}
