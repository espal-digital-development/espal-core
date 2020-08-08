// Code generated by espal-store-synthesizer. DO NOT EDIT.
package group

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/database/queryhelper"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/juju/errors"
)

var _ Store = &GroupsStore{}

// Store represents a data interaction object.
type Store interface {
	GetOneByID(id string) (*Group, bool, error)
	GetOneByIDWithCreator(id string) (*Group, bool, error)
	Delete(ids []string) error
	DeleteTranslation(ids []string) error
	ToggleActive(ids []string) error
	SetUserRights(id string, userRightIDs []string) error
	Name(userGroup *Group, languageID uint16) string
	TranslationsForID(userGroupID string) (translations []*Translation, ok bool, err error)
	Filter(context filters.QueryReader, language language) (result []*Group,
		filter filters.Filter, err error)
}

func (s *GroupsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Group, ok bool, err error) {
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
	result = make([]*Group, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		g := newGroup()
		fields := []interface{}{&g.id, &g.createdByID, &g.updatedByID, &g.createdAt, &g.updatedAt, &g.active, &g.userRights, &g.currencies}
		if withCreators {
			fields = append(fields, &g.createdByFirstName, &g.createdBySurname, &g.updatedByFirstName, &g.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, g)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of GroupsStore.
func New(selecterDatabase database.Database, updaterDatabase database.Database, deletorDatabase database.Database, databaseQueryHelper queryhelper.Helper, databaseFiltersFactory filters.Factory, translationsRepository translations.Repository, loggerService logger.Loggable) (*GroupsStore, error) {
	s := &GroupsStore{
		selecterDatabase:       selecterDatabase,
		updaterDatabase:        updaterDatabase,
		deletorDatabase:        deletorDatabase,
		databaseQueryHelper:    databaseQueryHelper,
		databaseFiltersFactory: databaseFiltersFactory,
		translationsRepository: translationsRepository,
		loggerService:          loggerService,
	}
	return s, nil
}
