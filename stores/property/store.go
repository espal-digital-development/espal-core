package property

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PropertiesStore data store.
type PropertiesStore struct {
	selecterDatabase database.Database
}
