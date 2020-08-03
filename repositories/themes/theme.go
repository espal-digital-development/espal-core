package themes

import (
	"sync"

	"github.com/juju/errors"
)

var _ Themeable = &Theme{}

// Themeable represents an object that provides visual view processing.
type Themeable interface {
	Code() string
	ParentCode() string
	SetView(view Viewable) error
	GetView(code string) (Viewable, error)
}

// Themes manages visual views.
type Theme struct {
	code       string
	parentCode string
	views      map[string]Viewable
	mutex      *sync.RWMutex
}

// Code returns the unique Theme code.
func (t *Theme) Code() string {
	return t.code
}

// ParentID returns the parent Theme unique identifier.
func (t *Theme) ParentCode() string {
	return t.parentCode
}

// SetView sets the given view into the internal Theme view-stack.
func (t *Theme) SetView(view Viewable) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for k := range t.views {
		if t.views[k].Code() == view.Code() {
			return errors.Errorf("view `%s` is already registered", view.Code())
		}
	}
	t.views[view.Code()] = view
	return nil
}

// TODO :: 777777 These 2 getters should try to look up the parent hierarchy to see if there's a parent View for the
// ID/Code.

// GetView returns the view for the given registered code.
func (t *Theme) GetView(code string) (Viewable, error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	if _, ok := t.views[code]; !ok {
		return nil, errors.Errorf("no view found with code `%s`", code)
	}
	return t.views[code], nil
}
