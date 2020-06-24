package personalmessage

import (
	"github.com/espal-digital-development/espal-core/database"
)

// PersonalMessagesStore data store.
type PersonalMessagesStore struct {
	selecterDatabase database.Database
}
