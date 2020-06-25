package newsarticle

import (
	"time"
)

// SectionTranslation database object.
// @synthesize
type SectionTranslation struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	language           uint16
	field              uint16
	value              string
	sectionID          string
}

// TableName returns the table name that belongs to the current model.
func (s *SectionTranslation) TableName() string {
	return "NewsArticleSectionTranslation"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *SectionTranslation) TableAlias() string {
	return "nast"
}
