package note

import (
	"github.com/espal-digital-development/espal-core/database"
)

// NotesStore data store.
type NotesStore struct {
	selecterDatabase database.Database
}
