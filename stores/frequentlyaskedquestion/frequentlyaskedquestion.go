package frequentlyaskedquestion

import (
	"time"
)

// FrequentlyAskedQuestion database object.
// @synthesize
type FrequentlyAskedQuestion struct {
	id                               string
	createdByID                      string
	updatedByID                      *string
	createdAt                        time.Time
	updatedAt                        *time.Time
	createdByFirstName               *string
	createdBySurname                 *string
	updatedByFirstName               *string
	updatedBySurname                 *string
	frequentlyAskedQuestionSectionID *string
	domainID                         string
	active                           bool
	sorting                          uint
}
