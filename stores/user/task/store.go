package task

import (
	"github.com/espal-digital-development/espal-core/database"
)

// TasksStore data store.
type TasksStore struct {
	selecterDatabase database.Database
}
