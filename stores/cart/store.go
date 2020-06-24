package cart

import (
	"github.com/espal-digital-development/espal-core/database"
)

// CartsStore data store.
type CartsStore struct {
	selecterDatabase database.Database
}
