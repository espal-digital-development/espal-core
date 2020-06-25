package store

import (
	"time"
)

// nolint:golint
// StoreEntity database object.
// @synthesize
type StoreEntity struct {
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
func (s *StoreEntity) TableName() string {
	return "Store"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *StoreEntity) TableAlias() string {
	return "sto"
}
