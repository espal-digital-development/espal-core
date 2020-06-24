package saleoffer

import (
	"github.com/espal-digital-development/espal-core/database"
)

// SaleOffersStore data store.
type SaleOffersStore struct {
	selecterDatabase database.Database
}
