package filters

import (
	"github.com/espal-digital-development/espal-core/database"
)

const defaultLimit = 5

var _ Factory = &Filters{}

// Factory represents an object that can Filters for database queries.
type Factory interface {
	NewFilter(queryReader QueryReader, m Model) Filter
}

// Filters acts as a spawner of Filter objects for database queries.
type Filters struct {
	selecterDatabase database.Database
}

// NewFilter returns a new ResultFilter based on the given model.
func (f *Filters) NewFilter(queryReader QueryReader, m Model) Filter {
	return &filter{
		queryReader:      queryReader,
		selecterDatabase: f.selecterDatabase,
		table:            m.TableName(),
		tableAlias:       m.TableAlias(),
		limit:            defaultLimit,
	}
}

// New returns a new instance of Filters.
func New(selecterDatabase database.Database) *Filters {
	return &Filters{
		selecterDatabase: selecterDatabase,
	}
}
