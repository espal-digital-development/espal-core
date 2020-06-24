package supplier

import (
	"time"
)

// Supplier database object.
// @synthesize
type Supplier struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	active             bool
	key                *string
	name               string
	contactFirstName   *string
	contactSurname     *string
	street             *string
	streetLine2        *string
	number             *string
	numberAddition     *string
	zipCode            *string
	city               *string
	state              *string
	country            *uint16
	phoneNumber        *string
	email              *string
	comments           *string

	// PurchaseOrders []PurchaseOrder
}

// TableAlias returns the unique resolved table alias for use in queries.
func (supplier *Supplier) TableAlias() string {
	return "sup"
}
