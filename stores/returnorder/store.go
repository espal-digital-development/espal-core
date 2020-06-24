package returnorder

import (
	"github.com/espal-digital-development/espal-core/database"
)

// ReturnOrdersStore data store.
type ReturnOrdersStore struct {
	selecterDatabase database.Database
}
