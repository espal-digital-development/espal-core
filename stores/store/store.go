package store

import (
	"github.com/espal-digital-development/espal-core/database"
)

// StoresStore data store.
type StoresStore struct {
	selecterDatabase database.Database
}
