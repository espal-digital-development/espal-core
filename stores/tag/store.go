package tag

import (
	"github.com/espal-digital-development/espal-core/database"
)

// TagsStore data store.
type TagsStore struct {
	selecterDatabase database.Database
}
