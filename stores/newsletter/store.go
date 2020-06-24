package newsletter

import (
	"github.com/espal-digital-development/espal-core/database"
)

// NewslettersStore data store.
type NewslettersStore struct {
	selecterDatabase database.Database
}
