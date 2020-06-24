package block

import (
	"github.com/espal-digital-development/espal-core/database"
)

// BlocksStore data store.
type BlocksStore struct {
	selecterDatabase database.Database
}
