package frequentlyaskedquestion

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
	parentID           *string // FrequentlyAskedQuestionSection
	domainID           string
	active             bool
	sorting            uint

	// Children []FrequentlyAskedQuestionSection
}

// TableName returns the table name that belongs to the current model.
func (section *Section) TableName() string {
	return "FrequentlyAskedQuestionSection"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (section *Section) TableAlias() string {
	return "faqs"
}
