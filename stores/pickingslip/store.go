package pickingslip

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PickingSlipsStore data store.
type PickingSlipsStore struct {
	selecterDatabase database.Database
}
