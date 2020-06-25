package paymentmethod

import (
	"time"
)

// PaymentMethod database object.
// @synthesize
type PaymentMethod struct {
	// TODO :: Images per domain/site/language
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	name               string
	description        *string
}

// TableAlias returns the unique resolved table alias for use in queries.
func (p *PaymentMethod) TableAlias() string {
	return "pmet"
}
