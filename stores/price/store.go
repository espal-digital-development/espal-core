package price

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PricesStore data store.
type PricesStore struct {
	selecterDatabase database.Database
}
