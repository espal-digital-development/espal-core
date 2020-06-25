package newsarticle

import (
	"time"
)

// NewsArticle database object.
// @synthesize
type NewsArticle struct {
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
	sectionID          *string // Section
	approvedByID       *string
	approvedDate       *time.Time
	publishDate        *time.Time
	expirationDate     *time.Time
	comments           *string

	// Likes []Like
}

// TableName returns the table name that belongs to the current model.
func (n *NewsArticle) TableName() string {
	return "NewsArticle"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (n *NewsArticle) TableAlias() string {
	return "na"
}
