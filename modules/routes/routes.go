package routes

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

// Handler represents an object that is able to handle a core context request.
type Handler interface {
	Handle(contexts.Context)
}

// Config Routes provider object.
type Config struct {
	Entries map[string]Handler
}

// Routes registration object.
type Routes struct {
	entries map[string]Handler
}

// Iterate iterates overall route entries and returns the path and handler
// per callback cycle.
func (r *Routes) Iterate(f func(path string, h Handler) error) error {
	if r.entries == nil {
		return nil
	}
	for path := range r.entries {
		if err := f(path, r.entries[path]); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// New returns a new instance of Routes.
func New(config *Config) (*Routes, error) {
	r := &Routes{
		entries: config.Entries,
	}
	return r, nil
}
