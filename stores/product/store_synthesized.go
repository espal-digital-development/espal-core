// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &ProductsStore{}

// Store represents a data interaction object.
type Store interface {
}

func (productsStore *ProductsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Model, ok bool, err error) {
	rows, err := productsStore.selecterDatabase.Query(query, params...)
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
	result = make([]*Model, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		model := newModel()
		fields := []interface{}{&model.id, &model.createdByID, &model.updatedByID, &model.createdAt, &model.updatedAt, &model.active, &model.sorting, &model.key, &model.taxGroupID, &model.nameRepresentationID, &model.descriptionRepresentationID, &model.imageRepresentationID}
		if withCreators {
			fields = append(fields, &model.createdByFirstName, &model.createdBySurname, &model.updatedByFirstName, &model.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, model)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of ProductsStore.
func New(selecterDatabase database.Database) (*ProductsStore, error) {
	productsStore := &ProductsStore{
		selecterDatabase: selecterDatabase,
	}
	return productsStore, nil
}
