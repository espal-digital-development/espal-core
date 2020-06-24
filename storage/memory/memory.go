package memory

import (
	"sync"

	"github.com/espal-digital-development/espal-core/storage"
)

var _ storage.Storage = &Storage{}

// Storage storage engine to interact with memory-stored objects.
type Storage struct {
	entries map[string][]byte
	mutex   *sync.RWMutex
}

// Exists returns an indicator if an entry exists for the given key.
func (storage *Storage) Exists(key string) bool {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()
	_, ok := storage.entries[key]
	return ok
}

// Set stores the value bytes at the given key.
func (storage *Storage) Set(key string, value []byte) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.entries[key] = value
	return nil
}

// Get fetches the stored bytes for the given key.
func (storage *Storage) Get(key string) ([]byte, bool, error) {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()
	value, ok := storage.entries[key]
	return value, ok, nil
}

// Delete wipes the bytes for the given key.
func (storage *Storage) Delete(key string) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	delete(storage.entries, key)
	return nil
}

// Iterate gives the possiblity to iterate over all entries.
// For the memory engine this can cause longer locks so don't use on heavy actions.
func (storage *Storage) Iterate(iterator func(key string, value []byte, err error) (keepCycling bool)) {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()
	for key, value := range storage.entries {
		if !iterator(key, value, nil) {
			break
		}
	}
}

// New returns a new instance of Storage that lives in memory.
func New() *Storage {
	return &Storage{
		mutex: &sync.RWMutex{},
	}
}
