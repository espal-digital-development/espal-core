package themes

import (
	"sync"

	"github.com/juju/errors"
)

var _ Repository = &Themes{}

// Repository represents a Themes repository.
type Repository interface {
	NewTheme(code string) *Theme
	Register(theme Themeable) error
	ThemeForID(id int) Themeable
}

// Themes manages the repository of visual themes.
type Themes struct {
	idTicker int
	entries  map[int]Themeable
	mutex    *sync.RWMutex
}

// Register registers the given Theme in the repository.
func (t *Themes) Register(theme Themeable) error {
	if theme.ID() <= 0 {
		return errors.Errorf("theme has no valid ID")
	}
	if theme.Code() == "" {
		return errors.Errorf("theme code cannot be empty")
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	var parentFound bool
	for k := range t.entries {
		if theme.ParentID() > 0 && theme.ParentID() == t.entries[k].ID() {
			parentFound = true
		}
		if t.entries[k].ID() == theme.ID() || t.entries[k].Code() == theme.Code() {
			return errors.Errorf("theme (ID: `%d`, Code: `%s`) already registered", theme.ID(), theme.Code())
		}
	}
	if theme.ParentID() > 0 && !parentFound {
		return errors.Errorf("theme parent ID `%d` does not exist", theme.ParentID())
	}
	t.entries[theme.ID()] = theme
	return nil
}

// ThemeForID returns the theme for the given registered id.
func (t *Themes) ThemeForID(id int) Themeable {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.entries[id]
}

// NewTheme returns a new instance of Theme.
func (t *Themes) NewTheme(code string) *Theme {
	t.idTicker++
	theme := &Theme{
		id:          t.idTicker,
		code:        code,
		views:       map[int]Viewable{},
		viewsByCode: map[string]Viewable{},
		mutex:       &sync.RWMutex{},
	}
	return theme
}

// New returns a new instance of Themes.
func New() (*Themes, error) {
	t := &Themes{
		entries: map[int]Themeable{},
		mutex:   &sync.RWMutex{},
	}
	return t, nil
}
