package slug

import (
	"strings"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

// SlugsStore data store.
type SlugsStore struct {
	selecterDatabase database.Database
}

// GetOneByDomainIDAndPath fetches by DomainID and Path.
func (s *SlugsStore) GetOneByDomainIDAndPath(domainID string, path string) (*Slug, bool, error) {
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	result, ok, err := s.fetch(`SELECT * FROM "Slug" WHERE "domainID" = $1 AND "path" = $2 LIMIT 1`, false, domainID, path)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}
