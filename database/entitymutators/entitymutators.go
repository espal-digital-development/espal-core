package entitymutators

import (
	"github.com/espal-digital-development/espal-core/database"
)

var _ Factory = &EntityMutators{}

// Factory represents an object that can spawn
// new EntityMutator instances.
type Factory interface {
	NewMutation(entity entity, form form, path string) Mutator
}

// EntityMutators is an assistant service to help running insert and
// and update calls for dynamic data in routes.
type EntityMutators struct {
	inserterDatabase database.Database
	updaterDatabase  database.Database
}

// NewMutation returns a new instance of Mutation to manipulate database data with.
func (m *EntityMutators) NewMutation(entity entity, form form, path string) Mutator {
	return &EntityMutator{
		inserterDatabase: m.inserterDatabase,
		updaterDatabase:  m.updaterDatabase,
		entity:           entity,
		formAction:       form.FieldValue("action"),
		path:             path,
		fields:           make([]string, 0),
		values:           make([]interface{}, 0),
	}
}

// New returns a new instance of EntityMutator.
func New(inserterDatabase database.Database, updaterDatabase database.Database) *EntityMutators {
	return &EntityMutators{
		inserterDatabase: inserterDatabase,
		updaterDatabase:  updaterDatabase,
	}
}
