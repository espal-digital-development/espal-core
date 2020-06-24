package download

import (
	"github.com/espal-digital-development/espal-core/database"
)

// DownloadsStore data store.
type DownloadsStore struct {
	selecterDatabase database.Database
}
