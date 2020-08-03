package themes

import "sync"

// DataStore represents a Setter/Getter storage object.
type DataStore interface {
	Set(k string, v interface{})
	Get(k string) (interface{}, bool)
}

// ViewData transferable View data communication object.
type ViewData struct {
	entries map[string]interface{}
	mutex   *sync.RWMutex
}

// Set sets the given value for the given key.
func (d *ViewData) Set(key string, value interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.entries[key] = value
}

// Get gets the value for the requested key.
func (d *ViewData) Get(key string) (interface{}, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	if v, ok := d.entries[key]; ok {
		return v, true
	}
	return nil, false
}

// NewViewData returns a new instance of ViewData.
func NewViewData() *ViewData {
	return &ViewData{
		entries: map[string]interface{}{},
		mutex:   &sync.RWMutex{},
	}
}
