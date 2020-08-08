// Code generated by espal-store-synthesizer. DO NOT EDIT.
package deliverymethod

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &DeliveryMethodsStore{}

// Store represents a data interaction object.
type Store interface {
}

func (s *DeliveryMethodsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*DeliveryMethod, ok bool, err error) {
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
	result = make([]*DeliveryMethod, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		d := newDeliveryMethod()
		fields := []interface{}{&d.id, &d.createdByID, &d.updatedByID, &d.createdAt, &d.updatedAt, &d.price}
		if withCreators {
			fields = append(fields, &d.createdByFirstName, &d.createdBySurname, &d.updatedByFirstName, &d.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, d)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of DeliveryMethodsStore.
func New(selecterDatabase database.Database) (*DeliveryMethodsStore, error) {
	s := &DeliveryMethodsStore{
		selecterDatabase: selecterDatabase,
	}
	return s, nil
}
