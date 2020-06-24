package paymentmethod

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PaymentMethodsStore data store.
type PaymentMethodsStore struct {
	selecterDatabase database.Database
}
