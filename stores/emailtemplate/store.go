package emailtemplate

import (
	"github.com/espal-digital-development/espal-core/database"
)

// EmailTemplatesStore data store.
type EmailTemplatesStore struct {
	selecterDatabase database.Database
}
