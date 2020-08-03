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
	GetTheme(code string) (Themeable, error)
}

// Themes manages the repository of visual themes.
type Themes struct {
	entries map[string]Themeable
	mutex   *sync.RWMutex
}

// Register registers the given Theme in the repository.
func (t *Themes) Register(theme Themeable) error {
	if theme.Code() == "" {
		return errors.Errorf("theme code cannot be empty")
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if theme.ParentCode() != "" {
		var parentFound bool
		for k := range t.entries {
			if theme.ParentCode() == t.entries[k].Code() {
				parentFound = true
			}
			if t.entries[k].Code() == theme.Code() {
				return errors.Errorf("theme `%s` is already registered", theme.Code())
			}
		}
		if !parentFound {
			return errors.Errorf("parent theme `%s` does not exist", theme.ParentCode())
		}
	}
	t.entries[theme.Code()] = theme
	return nil
}

// GetTheme returns the theme for the given registered code.
func (t *Themes) GetTheme(code string) (Themeable, error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	if _, ok := t.entries[code]; !ok {
		return nil, errors.Errorf("no theme found with code `%s`", code)
	}
	return t.entries[code], nil
}

// NewTheme returns a new instance of Theme.
func (t *Themes) NewTheme(code string) *Theme {
	theme := &Theme{
		code:  code,
		views: map[string]Viewable{},
		mutex: &sync.RWMutex{},
	}
	return theme
}

// New returns a new instance of Themes.
func New() (*Themes, error) {
	t := &Themes{
		entries: map[string]Themeable{},
		mutex:   &sync.RWMutex{},
	}
	return t, nil
}
