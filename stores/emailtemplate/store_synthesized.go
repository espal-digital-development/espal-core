// Code generated by espal-store-synthesizer. DO NOT EDIT.
package emailtemplate

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &EmailTemplatesStore{}

// Store represents a data interaction object.
type Store interface {
}

func (s *EmailTemplatesStore) fetch(query string, withCreators bool, params ...interface{}) (result []*EmailTemplate, ok bool, err error) {
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
	result = make([]*EmailTemplate, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		e := newEmailTemplate()
		fields := []interface{}{&e.id, &e.createdByID, &e.updatedByID, &e.createdAt, &e.updatedAt, &e.domainID, &e.active}
		if withCreators {
			fields = append(fields, &e.createdByFirstName, &e.createdBySurname, &e.updatedByFirstName, &e.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, e)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of EmailTemplatesStore.
func New(selecterDatabase database.Database) (*EmailTemplatesStore, error) {
	s := &EmailTemplatesStore{
		selecterDatabase: selecterDatabase,
	}
	return s, nil
}
