package receiving

import (
	"time"
)

// Receiving database object.
// @synthesize
type Receiving struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	comments           *string
}

// TableName returns the table name that belongs to the current model.
func (receiving *Receiving) TableName() string {
	return "PurchaseOrderReceiving"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (receiving *Receiving) TableAlias() string {
	return "porr"
}
