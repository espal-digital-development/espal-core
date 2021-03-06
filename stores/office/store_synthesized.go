// Code generated by espal-store-synthesizer. DO NOT EDIT.
package office

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &OfficesStore{}

// Store represents a data interaction object.
type Store interface {
}

func (s *OfficesStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Office, ok bool, err error) {
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
	result = make([]*Office, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		o := newOffice()
		fields := []interface{}{&o.id, &o.createdByID, &o.updatedByID, &o.createdAt, &o.updatedAt, &o.active, &o.sorting, &o.primaryContactPerson, &o.street, &o.streetLine2, &o.number, &o.numberAddition, &o.zipCode, &o.city, &o.state, &o.country, &o.phoneNumber, &o.email}
		if withCreators {
			fields = append(fields, &o.createdByFirstName, &o.createdBySurname, &o.updatedByFirstName, &o.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, o)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of OfficesStore.
func New(selecterDatabase database.Database) (*OfficesStore, error) {
	s := &OfficesStore{
		selecterDatabase: selecterDatabase,
	}
	return s, nil
}
