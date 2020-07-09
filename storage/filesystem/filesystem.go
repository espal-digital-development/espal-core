package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
	zglob "github.com/mattn/go-zglob"
)

var _ storage.Storage = &FileSystem{}

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
func (s *FileSystem) Iterate(iterator func(key string, value []byte, err error) (keepCycling bool)) {
	if s.path == "" {
		return
	}
	files, err := zglob.Glob(fmt.Sprintf("%s/**/*", s.path))
	if err != nil {
		iterator("", nil, errors.Trace(err))
		return
	}
	for k := range files {
		stat, err := os.Stat(files[k])
		if err != nil {
			if !iterator("", nil, errors.Trace(err)) {
				break
			}
		}
		// No features for directories. If ever need one; it can hook in here
		if stat.IsDir() {
			continue
		}
		fileBytes, err := ioutil.ReadFile(files[k])
		if !iterator(strings.Replace(files[k], s.pathWithoutRelativePrefix, "", 1), fileBytes, errors.Trace(err)) {
			break
		}
	}
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
