package user

import (
	"strings"
	"time"
)

// nolint:deadcode,unused
type userMethods interface {
	CurrenciesCount() uint
}

// User database object.
// @synthesize
type User struct {
	id                       string
	createdByID              string
	updatedByID              *string
	defaultDeliveryAddressID *string
	defaultInvoiceAddressID  *string
	createdAt                time.Time
	updatedAt                *time.Time
	createdByFirstName       *string
	createdBySurname         *string
	updatedByFirstName       *string
	updatedBySurname         *string
	active                   bool
	country                  *uint16
	language                 uint16
	firstName                *string
	surname                  *string
	dateOfBirth              *time.Time
	email                    string
	password                 string
	avatar                   *string
	priority                 uint // Can be used for everything concerning who comes first (like shipments or rewards)
	activationHash           *string
	activatedAt              *time.Time
	passwordResetHash        *string
	passwordResetLastSendAt  *time.Time
	passwordLastResetAt      *time.Time
	passwordResetCount       *uint8
	biography                *string
	comments                 *string // Back-end notes on the User
	currencies               string

	// UserGroupsUsers  []UserGroupUser
	// Addresses        []UserAddress
	// Contacts         []UserContact
	// Notes            []UserNote
	// PersonalMessages []PersonalMessage
}

// CurrenciesCount returns the amount of currencies that are defined in the string field.
func (u *User) CurrenciesCount() uint {
	if u.currencies != "" {
		return uint(strings.Count(u.currencies, ",") + 1)
	}
	return 0
}
