package paymenttransaction

import (
	"time"
)

// PaymentTransaction database object.
// @synthesize
type PaymentTransaction struct {
	// TODO :: Make this object compatible to hold all fields that it's
	// linked PaymentMethod requires.
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	paymentAccountID   string
	saleOrderID        string
	responseCode       int
	amount             float32
	hash               *string
	message            *string
}
