// Code generated by espal-store-synthesizer. DO NOT EDIT.
package menu

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &MenusStore{}

// Store represents a data interaction object.
type Store interface {
}

func (s *MenusStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Menu, ok bool, err error) {
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
	result = make([]*Menu, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		m := newMenu()
		fields := []interface{}{&m.id, &m.createdByID, &m.updatedByID, &m.createdAt, &m.updatedAt, &m.active, &m.sorting, &m.slugID, &m.externalLink, &m.internalLink, &m.parentID}
		if withCreators {
			fields = append(fields, &m.createdByFirstName, &m.createdBySurname, &m.updatedByFirstName, &m.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, m)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of MenusStore.
func New(selecterDatabase database.Database) (*MenusStore, error) {
	s := &MenusStore{
		selecterDatabase: selecterDatabase,
	}
	return s, nil
}
