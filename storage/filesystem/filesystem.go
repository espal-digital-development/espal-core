package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
)

var _ storage.Storage = &FileSystem{}

var (
	errNotSupported = errors.New("this feature is not supported for filesystems")
	errStopCycling  = errors.New("stopped cycling")
)

// TODO :: 7777 Any Windows path issues can now be funneled as all FS interaction should only happen here

// FileSystem storage engine to interact with a local filesystem.
type FileSystem struct {
	path                      string
	pathWithoutRelativePrefix string
}

// Exists returns an indicator if an entry exists for the given key.
func (s *FileSystem) Exists(key string) bool {
	if _, err := os.Stat(s.path + key); !os.IsNotExist(err) {
		return true
	}
	return false
}

// Set stores the value bytes at the given key.
func (s *FileSystem) Set(key string, value []byte) error {
	// TODO :: 777 Create all intermediate directories if needed
	return errors.Trace(ioutil.WriteFile(s.path+key, value, 0600))
}

// Get fetches the stored bytes for the given key.
func (s *FileSystem) Get(key string) ([]byte, bool, error) {
	if _, err := os.Stat(s.path + key); os.IsNotExist(err) {
		return nil, false, nil
	}
	read, err := ioutil.ReadFile(s.path + key)
	if err != nil {
		return nil, false, errors.Trace(err)
	}
	return read, true, nil
}

// Delete wipes the bytes for the given key.
func (s *FileSystem) Delete(key string) error {
	return errors.Trace(os.Remove(s.path + key))
}

// Iterate gives the possiblity to iterate over all entries.
func (s *FileSystem) Iterate(iterator func(key string, value []byte, err error) (keepCycling bool)) error {
	if s.path == "" {
		return nil
	}
	err := filepath.Walk(s.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Trace(err)
		}
		// No features for directories. If ever need one; it can hook in here
		if info.IsDir() {
			return nil
		}
		fileBytes, err := ioutil.ReadFile(path)
		if !iterator(strings.Replace(path, s.pathWithoutRelativePrefix, "", 1), fileBytes, errors.Trace(err)) {
			return errors.Trace(errStopCycling)
		}
		return nil
	})
	if err != nil && err != errStopCycling {
		return errors.Trace(err)
	}
	return nil
}

// LoadAllFromPath walks all the files in the given path and it's subdirectories
// and loads it into the storage.
func (s *FileSystem) LoadAllFromPath(subjectPath string) error {
	return errors.Trace(errNotSupported)
}

// New returns a new instance of FileSystem that interacts with the filesystem.
func New(path string) (*FileSystem, error) {
	f := &FileSystem{}
	if path != "" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil, errors.Trace(err)
		}
		f = &FileSystem{
			path: strings.TrimSuffix(path, "/") + "/",
		}
		f.pathWithoutRelativePrefix = strings.TrimPrefix(f.path, "./")
	}
	return f, nil
}
