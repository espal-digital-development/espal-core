package poll

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PollsStore data store.
type PollsStore struct {
	selecterDatabase database.Database
}
