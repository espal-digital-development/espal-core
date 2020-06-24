package product

import (
	"github.com/espal-digital-development/espal-core/database"
)

// ProductsStore data store.
type ProductsStore struct {
	selecterDatabase database.Database
}
