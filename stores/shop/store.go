package shop

import (
	"github.com/espal-digital-development/espal-core/database"
)

// ShopsStore data store.
type ShopsStore struct {
	selecterDatabase database.Database
}
