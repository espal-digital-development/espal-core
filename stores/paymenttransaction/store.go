package paymenttransaction

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PaymentTransactionsStore data store.
type PaymentTransactionsStore struct {
	selecterDatabase database.Database
}
