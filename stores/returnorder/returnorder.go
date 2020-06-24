package returnorder

import (
	"time"
)

// ReturnOrder database object.
// @synthesize
type ReturnOrder struct {
	// SaleOrder isn't linked on this level, as someone might aswel return
	// multiple OrderLines of different Orders, and in different quantities.
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	domainID           string
	userID             string
	comments           *string

	// Lines []ReturnOrderLine
}
