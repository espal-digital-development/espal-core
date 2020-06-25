package address

import (
	"time"
)

// Address database object.
// @synthesize
type Address struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	userID             string
	active             bool
	firstName          *string
	surname            *string
	street             string
	streetLine2        *string
	number             string
	numberAddition     *string
	zipCode            string
	city               string
	state              *string
	country            *uint16
	phoneNumber        *string
	email              *string
	// TODO :: Add sorting?
}

// TableName returns the table name that belongs to the current model.
func (a *Address) TableName() string {
	return "UserAddress"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (a *Address) TableAlias() string {
	return "ua"
}
