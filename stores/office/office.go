package office

import (
	"time"
)

// Office database object.
// @synthesize
type Office struct {
	id                   string
	createdByID          string
	updatedByID          *string
	createdAt            time.Time
	updatedAt            *time.Time
	createdByFirstName   *string
	createdBySurname     *string
	updatedByFirstName   *string
	updatedBySurname     *string
	active               bool
	sorting              uint
	primaryContactPerson *string
	street               string
	streetLine2          *string
	number               string
	numberAddition       *string
	zipCode              string
	city                 string
	state                *string
	country              *uint16
	phoneNumber          *string
	email                *string
}

// TableName returns the table name that belongs to the current model.
func (o *Office) TableName() string {
	return "Office"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (o *Office) TableAlias() string {
	return "o"
}
