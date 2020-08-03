package themes

import (
	"github.com/espal-digital-development/espal-core/repositories/themes"
)

// Themeable represents a Theme and View provider.
type Themeable interface {
	Themes() []themes.Themeable
	Views() map[string][]themes.Viewable
}

// Config Themes provider object.
type Config struct {
	Entries []themes.Themeable
	Views   map[string][]themes.Viewable
}

// Themes provider object.
type Themes struct {
	entries []themes.Themeable
	views   map[string][]themes.Viewable
}

// Themes returns the registered themes.
func (t *Themes) Themes() []themes.Themeable {
	return t.entries
}

// Views returns the registered dedicated views for specific themes.
func (t *Themes) Views() map[string][]themes.Viewable {
	return t.views
}

// New returns a new instance of Themes.
func New(config *Config) (*Themes, error) {
	r := &Themes{
		entries: config.Entries,
		views:   config.Views,
	}
	return r, nil
}
