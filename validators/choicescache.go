package validators

import (
	"sync"
)

type choicesCache struct {
	choices map[string]options
	mutex   *sync.RWMutex
}

func newChoicesCache() *choicesCache {
	return &choicesCache{
		choices: map[string]options{},
		mutex:   &sync.RWMutex{},
	}
}

// Store the CountryChoice to the given key.
func (c *choicesCache) Store(key string, value options) {
	c.mutex.Lock()
	c.choices[key] = value
	c.mutex.Unlock()
}

// Load the CountryChoice based on the requested key.
func (c *choicesCache) Load(key string) (options, bool) {
	c.mutex.RLock()
	value, ok := c.choices[key]
	c.mutex.RUnlock()
	if !ok {
		return options{}, false
	}
	return value, true
}

// Consume the CountryChoice based on the requested key and destroy it afterwards.
func (c *choicesCache) Consume(key string) (options, bool) {
	c.mutex.RLock()
	value, ok := c.choices[key]
	c.mutex.RUnlock()
	if !ok {
		return options{}, false
	}
	c.mutex.Lock()
	delete(c.choices, key)
	c.mutex.Unlock()

	return value, true
}

// Count returns the current cache's entry-count.
func (c *choicesCache) Count() uint {
	return uint(len(c.choices))
}
