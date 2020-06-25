package reseller

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
	resellerID         string
}

// TableName returns the table name that belongs to the current model.
func (t *Translation) TableName() string {
	return "ResellerTranslation"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (t *Translation) TableAlias() string {
	return "rt"
}
