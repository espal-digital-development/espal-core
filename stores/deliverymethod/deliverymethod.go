package deliverymethod

import (
	"time"
)

// DeliveryMethod database object.
// @synthesize
type DeliveryMethod struct {
	// TODO :: Link Prices per currency per country/usergroup/user
	// to DeliveryMethod.
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	price              float64
}
