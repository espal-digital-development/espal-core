package themes

import (
	"sync"

	"github.com/juju/errors"
)

var _ Themeable = &Theme{}

// Themeable represents an object that provides visual view processing.
type Themeable interface {
	ID() int
	Code() string
	ParentID() int
	AddView(view Viewable) error
	ViewForID(id int) Viewable
	ViewForCode(code string) Viewable
}

// Themes manages visual views.
type Theme struct {
	id          int
	code        string
	parentID    int
	views       map[int]Viewable
	viewsByCode map[string]Viewable
	mutex       *sync.RWMutex
}

// ID returns the unique Theme identifier.
func (t *Theme) ID() int {
	return t.id
}

// Code returns the unique Theme code.
func (t *Theme) Code() string {
	return t.code
}

// ParentID returns the parent Theme unique identifier.
func (t *Theme) ParentID() int {
	return t.parentID
}

// AddView adds the given view into the internal Theme view-stack.
func (t *Theme) AddView(view Viewable) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for k := range t.views {
		if t.views[k].ID() == view.ID() || t.views[k].Code() == view.Code() {
			return errors.Errorf("view (ID: `%d`, Code: `%s`) already registered", view.ID(), view.Code())
		}
	}
	t.views[view.ID()] = view
	t.viewsByCode[view.Code()] = view
	return nil
}

// ViewForID returns the view for the given registered id.
func (t *Theme) ViewForID(id int) Viewable {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.views[id]
}

// ViewForCode returns the view for the given registered code.
func (t *Theme) ViewForCode(code string) Viewable {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.viewsByCode[code]
}
