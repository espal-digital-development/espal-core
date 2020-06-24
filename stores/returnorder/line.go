package returnorder

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	returnOrderID      string
	saleOrderLineID    string
	quantity           uint
	comments           *string
}

// TableName returns the table name that belongs to the current model.
func (line *Line) TableName() string {
	return "ReturnOrderLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (line *Line) TableAlias() string {
	return "rol"
}
