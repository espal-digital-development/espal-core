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
func (choicesCache *choicesCache) Store(key string, value options) {
	choicesCache.mutex.Lock()
	choicesCache.choices[key] = value
	choicesCache.mutex.Unlock()
}

// Load the CountryChoice based on the requested key.
func (choicesCache *choicesCache) Load(key string) (options, bool) {
	choicesCache.mutex.RLock()
	value, ok := choicesCache.choices[key]
	choicesCache.mutex.RUnlock()
	if !ok {
		return options{}, false
	}
	return value, true
}

// Consume the CountryChoice based on the requested key and destroy it afterwards.
func (choicesCache *choicesCache) Consume(key string) (options, bool) {
	choicesCache.mutex.RLock()
	value, ok := choicesCache.choices[key]
	choicesCache.mutex.RUnlock()
	if !ok {
		return options{}, false
	}
	choicesCache.mutex.Lock()
	delete(choicesCache.choices, key)
	choicesCache.mutex.Unlock()

	return value, true
}

// Count returns the current cache's entry-count.
func (choicesCache *choicesCache) Count() uint {
	return uint(len(choicesCache.choices))
}
