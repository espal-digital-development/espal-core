package tax

import (
	"github.com/espal-digital-development/espal-core/database"
)

// TaxesStore data store.
type TaxesStore struct {
	selecterDatabase database.Database
}
