package saleorder

import (
	"github.com/espal-digital-development/espal-core/database"
)

// SaleOrdersStore data store.
type SaleOrdersStore struct {
	selecterDatabase database.Database
}
