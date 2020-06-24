package shipment

import (
	"time"
)

// Cost database object.
// @synthesize
type Cost struct {
	// TODO :: Implement (per country?)
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
}

// TableName returns the table name that belongs to the current model.
func (cost *Cost) TableName() string {
	return "ShipmentCost"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (cost *Cost) TableAlias() string {
	return "shc"
}
