package deliverymethod

import (
	"github.com/espal-digital-development/espal-core/database"
)

// DeliveryMethodsStore data store.
type DeliveryMethodsStore struct {
	selecterDatabase database.Database
}
