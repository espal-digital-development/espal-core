// Code generated by espal-store-synthesizer. DO NOT EDIT.
package credit

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &CreditsStore{}

// Store represents a data interaction object.
type Store interface {
}

func (creditsStore *CreditsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Credit, ok bool, err error) {
	rows, err := creditsStore.selecterDatabase.Query(query, params...)
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
	result = make([]*Credit, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		credit := newCredit()
		fields := []interface{}{&credit.id, &credit.createdByID, &credit.updatedByID, &credit.createdAt, &credit.updatedAt}
		if withCreators {
			fields = append(fields, &credit.createdByFirstName, &credit.createdBySurname, &credit.updatedByFirstName, &credit.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, credit)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of CreditsStore.
func New(selecterDatabase database.Database) (*CreditsStore, error) {
	creditsStore := &CreditsStore{
		selecterDatabase: selecterDatabase,
	}
	return creditsStore, nil
}
