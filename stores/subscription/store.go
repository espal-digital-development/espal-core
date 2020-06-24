package subscription

import (
	"github.com/espal-digital-development/espal-core/database"
)

// SubscriptionsStore data store.
type SubscriptionsStore struct {
	selecterDatabase database.Database
}
