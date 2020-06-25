package product

import (
	"time"
)

// ModelDimension database object.
// @synthesize
type ModelDimension struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	sorting            uint
	key                *string
	modelID            string
	propertyID         string
}

// TableName returns the table name that belongs to the current model.
func (m *ModelDimension) TableName() string {
	return "ProductModelDimension"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (m *ModelDimension) TableAlias() string {
	return "promd"
}
