// Code generated by espal-store-synthesizer. DO NOT EDIT.
package cart

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &CartsStore{}

// Store represents a data interaction object.
type Store interface {
}

func (cartsStore *CartsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Cart, ok bool, err error) {
	rows, err := cartsStore.selecterDatabase.Query(query, params...)
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
	result = make([]*Cart, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		cart := newCart()
		fields := []interface{}{&cart.id, &cart.createdByID, &cart.updatedByID, &cart.createdAt, &cart.updatedAt, &cart.domainID, &cart.userID}
		if withCreators {
			fields = append(fields, &cart.createdByFirstName, &cart.createdBySurname, &cart.updatedByFirstName, &cart.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, cart)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of CartsStore.
func New(selecterDatabase database.Database) (*CartsStore, error) {
	cartsStore := &CartsStore{
		selecterDatabase: selecterDatabase,
	}
	return cartsStore, nil
}
