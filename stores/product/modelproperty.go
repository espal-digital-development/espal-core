package product

import (
	"time"
)

// ModelProperty database object.
// @synthesize
type ModelProperty struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	productModelID     string
	propertyID         string
	sorting            uint
	key                *string
}

// TableName returns the table name that belongs to the current model.
func (m *ModelProperty) TableName() string {
	return "ProductModelProperty"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (m *ModelProperty) TableAlias() string {
	return "promp"
}
