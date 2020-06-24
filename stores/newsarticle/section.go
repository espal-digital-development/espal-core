package newsarticle

import (
	"time"
)

// Section database object.
// @synthesize
type Section struct {
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
	parentID           *string // Section

	// Children []Section
}

// TableName returns the table name that belongs to the current model.
func (section *Section) TableName() string {
	return "NewsArticleSection"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (section *Section) TableAlias() string {
	return "nas"
}
