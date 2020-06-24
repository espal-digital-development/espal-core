// Code generated by espal-store-synthesizer. DO NOT EDIT.
package group

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &GroupsStore{}

// Store represents a data interaction object.
type Store interface {
}

func (groupsStore *GroupsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Group, ok bool, err error) {
	rows, err := groupsStore.selecterDatabase.Query(query, params...)
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
	result = make([]*Group, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		group := newGroup()
		fields := []interface{}{&group.id, &group.createdByID, &group.updatedByID, &group.createdAt, &group.updatedAt, &group.active, &group.sorting, &group.code}
		if withCreators {
			fields = append(fields, &group.createdByFirstName, &group.createdBySurname, &group.updatedByFirstName, &group.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, group)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of GroupsStore.
func New(selecterDatabase database.Database) (*GroupsStore, error) {
	groupsStore := &GroupsStore{
		selecterDatabase: selecterDatabase,
	}
	return groupsStore, nil
}
