package invoice

import (
	"github.com/espal-digital-development/espal-core/database"
)

// InvoicesStore data store.
type InvoicesStore struct {
	selecterDatabase database.Database
}
