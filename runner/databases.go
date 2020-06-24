package runner

import (
	"github.com/espal-digital-development/espal-core/database"
)

type databases struct {
	selecter database.Database
	creator  database.Database
	inserter database.Database
	updater  database.Database
	deletor  database.Database
}
