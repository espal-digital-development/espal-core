package memory

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
)

var _ storage.Storage = &Storage{}

// Storage storage engine to interact with memory-stored objects.
type Storage struct {
	entries map[string][]byte
	mutex   *sync.RWMutex
}

// Exists returns an indicator if an entry exists for the given key.
func (s *Storage) Exists(key string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	_, ok := s.entries[key]
	return ok
}

// Set stores the value bytes at the given key.
func (s *Storage) Set(key string, value []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.entries[key] = value
	return nil
}

// Get fetches the stored bytes for the given key.
func (s *Storage) Get(key string) ([]byte, bool, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	value, ok := s.entries[key]
	return value, ok, nil
}

// Delete wipes the bytes for the given key.
func (s *Storage) Delete(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.entries, key)
	return nil
}

// Iterate gives the possiblity to iterate over all entries.
// For the memory engine this can cause longer locks so don't use on heavy actions.
func (s *Storage) Iterate(iterator func(key string, value []byte, err error) (keepCycling bool)) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for key, value := range s.entries {
		if !iterator(key, value, nil) {
			break
		}
	}
	return nil
}

// LoadAllFromPath walks all the files in the given path and it's subdirectories and loads it into the storage.
func (s *Storage) LoadAllFromPath(subjectPath string) error {
	if subjectPath == "" {
		return nil
	}
	subjectPath, err := filepath.Abs(subjectPath)
	if err != nil {
		return errors.Trace(err)
	}
	return errors.Trace(filepath.Walk(subjectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Trace(err)
		}
		if info.IsDir() {
			return nil
		}
		fileBytes, fileErr := ioutil.ReadFile(path)
		if fileErr != nil {
			return errors.Trace(fileErr)
		}
		setErr := s.Set(strings.TrimPrefix(path, subjectPath+"/"),
			fileBytes)
		if setErr != nil {
			return errors.Trace(setErr)
		}
		return nil
	}))
}

// New returns a new instance of Storage that lives in memory.
func New() *Storage {
	return &Storage{
		entries: map[string][]byte{},
		mutex:   &sync.RWMutex{},
	}
}
