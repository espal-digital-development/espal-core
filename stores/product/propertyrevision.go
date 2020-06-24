package product

import (
	"time"
)

// PropertyRevision database object.
// @synthesize
type PropertyRevision struct {
	// TODO :: On PreSave/PreUpdate check if the field-combinations are valid
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	domainID           *string
	language           *uint16
	bundledID          *string
	modelID            *string
	variantID          *string
	propertyID         string
	propertyOptionID   *string
	revertedFromID     *string // PPRevision
	revision           uint
	value              *string
}

// TableName returns the table name that belongs to the current model.
func (propertyRevision *PropertyRevision) TableName() string {
	return "ProductPropertyRevision"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (propertyRevision *PropertyRevision) TableAlias() string {
	return "propr"
}
