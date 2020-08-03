package slug

import (
	"strings"
	"sync"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

// SlugsStore data store.
type SlugsStore struct {
	selecterDatabase database.Database

	cacheNormal map[string]map[string]*Slug
	mutex       *sync.RWMutex
}

// GetOneByDomainIDAndPath fetches by DomainID and Path.
func (s *SlugsStore) GetOneByDomainIDAndPath(domainID string, path string) (*Slug, bool, error) {
	// TODO :: 77777 Move this caching to the general cache notifier
	if s.mutex == nil {
		s.mutex = &sync.RWMutex{}
		s.cacheNormal = make(map[string]map[string]*Slug)
	}
	s.mutex.RLock()

	if v, ok := s.cacheNormal[domainID][path]; ok {
		s.mutex.RUnlock()
		return v, true, nil
	}
	s.mutex.RUnlock()

	finalPath := path
	if strings.HasPrefix(path, "/") {
		finalPath = strings.TrimPrefix(path, "/")
	}

	result, ok, err := s.fetch(`SELECT * FROM "Slug" WHERE "domainID" = $1 AND "path" = $2 LIMIT 1`, false, domainID,
		finalPath)
	if len(result) == 1 {
		s.mutex.Lock()
		s.cacheNormal[domainID][path] = result[0]
		s.mutex.Unlock()
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}
