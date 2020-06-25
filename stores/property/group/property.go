package group

import (
	"time"
)

// Property database object.
// @synthesize
type Property struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	propertyGroupID    string
	propertyID         string
}

// TableName returns the table name that belongs to the current model.
func (p *Property) TableName() string {
	return "PropertyGroupProperty"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (p *Property) TableAlias() string {
	return "prgrp"
}
