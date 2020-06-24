package property

import (
	"time"
)

// OptionTranslation database object.
// @synthesize
type OptionTranslation struct {
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
	optionID           string
}

// TableName returns the table name that belongs to the current model.
func (optionTranslation *OptionTranslation) TableName() string {
	return "PropertyOptionTranslation"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (optionTranslation *OptionTranslation) TableAlias() string {
	return "proot"
}
