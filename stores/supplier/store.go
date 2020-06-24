package supplier

import (
	"github.com/espal-digital-development/espal-core/database"
)

// SuppliersStore data store.
type SuppliersStore struct {
	selecterDatabase database.Database
}
