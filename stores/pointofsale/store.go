package pointofsale

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PointsOfSaleStore data store.
type PointsOfSaleStore struct {
	selecterDatabase database.Database
}
