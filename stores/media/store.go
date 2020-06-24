package media

import (
	"github.com/espal-digital-development/espal-core/database"
)

// MediasStore data store.
type MediasStore struct {
	// This struct doesn't have the best plural name, but as Media
	// is already plural for Medium let's keep it clumpsy like this.
	selecterDatabase database.Database
}
