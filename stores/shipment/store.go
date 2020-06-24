package shipment

import (
	"github.com/espal-digital-development/espal-core/database"
)

// ShipmentsStore data store.
type ShipmentsStore struct {
	selecterDatabase database.Database
}
