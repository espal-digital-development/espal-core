// Code generated by espal-store-synthesizer. DO NOT EDIT.
package slug

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &SlugsStore{}

// Store represents a data interaction object.
type Store interface {
	GetOneByDomainIDAndPath(domainID string, path string) (*Slug, bool, error)
}

func (s *SlugsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Slug, ok bool, err error) {
	rows, err := s.selecterDatabase.Query(query, params...)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		err = errors.Trace(err)
		return
	}
	defer func(dbRows database.Rows) {
		closeErr := dbRows.Close()
		if err != nil && closeErr != nil {
			err = errors.Wrap(err, closeErr)
		} else if closeErr != nil {
			err = errors.Trace(closeErr)
		}
	}(rows)
	result = make([]*Slug, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		s := newSlug()
		fields := []interface{}{&s.id, &s.createdByID, &s.updatedByID, &s.createdAt, &s.updatedAt, &s.domainID, &s.language, &s.path, &s.rerouteTo, &s.invalidWithStatus, &s.invalidMessage, &s.redirectToRawPath, &s.redirectStatusCode}
		if withCreators {
			fields = append(fields, &s.createdByFirstName, &s.createdBySurname, &s.updatedByFirstName, &s.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, s)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of SlugsStore.
func New(selecterDatabase database.Database) (*SlugsStore, error) {
	s := &SlugsStore{
		selecterDatabase: selecterDatabase,
	}
	return s, nil
}
