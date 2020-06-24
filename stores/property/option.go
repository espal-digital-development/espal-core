package property

import (
	"time"
)

// Option database object.
// @synthesize
type Option struct {
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
	key                *string
	propertyID         string
}

// TableName returns the table name that belongs to the current model.
func (option *Option) TableName() string {
	return "PropertyOption"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (option *Option) TableAlias() string {
	return "proo"
}
