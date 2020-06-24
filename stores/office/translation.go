package office

import (
	"time"
)

// Translation database object.
// @synthesize
type Translation struct {
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
	officeID           string
}

// TableName returns the table name that belongs to the current model.
func (translation *Translation) TableName() string {
	return "OfficeTranslation"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (translation *Translation) TableAlias() string {
	return "ot"
}
