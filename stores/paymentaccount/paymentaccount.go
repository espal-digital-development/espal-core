package paymentaccount

import (
	"time"
)

// PaymentAccount database object.
// @synthesize
type PaymentAccount struct {
	// TODO :: Make this object compatible to hold all fields that it's
	// linked PaymentMethod requires to validate the account.
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	paymentMethodID    string
	active             bool
	name               string
	username           *string
	passphrase         *string
	secretKey          *string
	publicKey          *string
	certificate        *string
}
