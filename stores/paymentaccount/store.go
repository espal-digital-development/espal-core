package paymentaccount

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PaymentAccountsStore data store.
type PaymentAccountsStore struct {
	selecterDatabase database.Database
}
