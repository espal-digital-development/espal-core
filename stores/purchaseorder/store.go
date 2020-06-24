package purchaseorder

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PurchaseOrdersStore data store.
type PurchaseOrdersStore struct {
	selecterDatabase database.Database
}
