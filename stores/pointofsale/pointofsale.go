package pointofsale

import (
	"time"
)

// PointOfSale database object.
// @synthesize
type PointOfSale struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	shopID             string
}
