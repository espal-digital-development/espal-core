package gift

import (
	"github.com/espal-digital-development/espal-core/database"
)

// GiftsStore data store.
type GiftsStore struct {
	selecterDatabase database.Database
}
