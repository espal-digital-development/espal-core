// Code generated by espal-store-synthesizer. DO NOT EDIT.
package frequentlyaskedquestion

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &FrequentlyAskedQuestionsStore{}

// Store represents a data interaction object.
type Store interface {
}

func (f *FrequentlyAskedQuestionsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*FrequentlyAskedQuestion, ok bool, err error) {
	rows, err := f.selecterDatabase.Query(query, params...)
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
	result = make([]*FrequentlyAskedQuestion, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		f := newFrequentlyAskedQuestion()
		fields := []interface{}{&f.id, &f.createdByID, &f.updatedByID, &f.createdAt, &f.updatedAt, &f.frequentlyAskedQuestionSectionID, &f.domainID, &f.active, &f.sorting}
		if withCreators {
			fields = append(fields, &f.createdByFirstName, &f.createdBySurname, &f.updatedByFirstName, &f.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, f)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of FrequentlyAskedQuestionsStore.
func New(selecterDatabase database.Database) (*FrequentlyAskedQuestionsStore, error) {
	f := &FrequentlyAskedQuestionsStore{
		selecterDatabase: selecterDatabase,
	}
	return f, nil
}
