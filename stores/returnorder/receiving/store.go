package receiving

import (
	"github.com/espal-digital-development/espal-core/database"
)

// ReceivingsStore data store.
type ReceivingsStore struct {
	selecterDatabase database.Database
}
