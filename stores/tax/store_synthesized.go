// Code generated by espal-store-synthesizer. DO NOT EDIT.
package tax

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &TaxesStore{}

// Store represents a data interaction object.
type Store interface {
}

func (taxesStore *TaxesStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Tax, ok bool, err error) {
	rows, err := taxesStore.selecterDatabase.Query(query, params...)
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
	result = make([]*Tax, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		tax := newTax()
		fields := []interface{}{&tax.id, &tax.createdByID, &tax.updatedByID, &tax.createdAt, &tax.updatedAt, &tax.taxGroupID, &tax.country, &tax.rate}
		if withCreators {
			fields = append(fields, &tax.createdByFirstName, &tax.createdBySurname, &tax.updatedByFirstName, &tax.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, tax)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of TaxesStore.
func New(selecterDatabase database.Database) (*TaxesStore, error) {
	taxesStore := &TaxesStore{
		selecterDatabase: selecterDatabase,
	}
	return taxesStore, nil
}
