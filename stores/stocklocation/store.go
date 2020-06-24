package stocklocation

import (
	"github.com/espal-digital-development/espal-core/database"
)

// StockLocationsStore data store.
type StockLocationsStore struct {
	selecterDatabase database.Database
}
