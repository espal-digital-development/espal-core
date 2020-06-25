package group

import (
	"database/sql"
	"strings"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/juju/errors"
)

type language interface {
	ID() uint16
}

// GroupsStore data store.
type GroupsStore struct {
	selecterDatabase       database.Database
	updaterDatabase        database.Database
	deletorDatabase        database.Database
	databaseFiltersFactory filters.Factory
	translationsRepository translations.Repository
	loggerService          logger.Loggable
}

// GetOneByID fetches by ID.
func (g *GroupsStore) GetOneByID(id string) (*Group, bool, error) {
	result, ok, err := g.fetch(`SELECT * FROM "UserGroup" WHERE "id" = $1 LIMIT 1`, false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (g *GroupsStore) GetOneByIDWithCreator(id string) (*Group, bool, error) {
	result, ok, err := g.fetch(`SELECT ug.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "UserGroup" ug
		LEFT JOIN "User" cu ON cu."id" = ug."createdByID"
		LEFT JOIN "User" uu ON uu."id" = ug."updatedByID"
		WHERE ug."id" = $1 LIMIT 1`, true, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// Delete deletes the given ID(s).
func (g *GroupsStore) Delete(ids []string) error {
	transaction, err := g.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(`DELETE FROM "UserGroupTranslation" WHERE "userGroupID" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(`DELETE FROM "UserGroup" WHERE "id" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// DeleteTranslation deletes the given translation ID(s).
func (g *GroupsStore) DeleteTranslation(ids []string) error {
	transaction, err := g.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(`DELETE FROM "UserGroupTranslation" WHERE "userGroupID" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// ToggleActive toggles the active state of the given ID(s).
func (g *GroupsStore) ToggleActive(ids []string) error {
	transaction, err := g.updaterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Query(`UPDATE "UserGroup" SET "active" = NOT "active" WHERE "id" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// SetUserRights sets the userRightIDs for the given id.
func (g *GroupsStore) SetUserRights(id string, userRightIDs []string) error {
	_, err := g.updaterDatabase.Exec(`UPDATE "UserGroup" SET "userRights" = $1 WHERE "id" = $1`, strings.Join(userRightIDs, ","), id)
	return errors.Trace(err)
}

// Name returns the presentable name.
func (g *GroupsStore) Name(userGroup *Group, languageID uint16) string {
	var name string
	err := g.selecterDatabase.QueryRow(
		`SELECT "value" FROM "UserGroupTranslation" WHERE "userGroupID" = $1 AND "field" = $2 AND "language" = $3`,
		userGroup.ID(), database.DBTranslationFieldName, languageID).Scan(&name)
	if err != nil && err != sql.ErrNoRows {
		g.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	if err == sql.ErrNoRows || name == "" {
		return g.translationsRepository.Singular(languageID, "userGroup") + " " + userGroup.ID()
	}
	return name
}

// TranslationsForID fetches UserGroupTranslations for userGroupID.
func (g *GroupsStore) TranslationsForID(userGroupID string) (translations []*Translation, ok bool, err error) {
	rows, err := g.selecterDatabase.Query(`SELECT ugt.*
		FROM "UserGroupTranslation" ugt
		LEFT JOIN "UserGroup" ug ON ugt."userGroupID" = ug."id"
		WHERE ugt."userGroupID" = $1`, userGroupID)
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
	translations = make([]*Translation, 0)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			err = errors.Trace(err)
			return
		}
		translation := newTranslation()
		err = rows.Scan(&translation.id, &translation.createdByID, &translation.updatedByID,
			&translation.groupID, &translation.createdAt, &translation.updatedAt,
			&translation.language, &translation.field, &translation.value)
		if err != nil {
			err = errors.Trace(err)
			return nil, false, errors.Trace(err)
		}
		translations = append(translations, translation)
	}

	if err = rows.Close(); err != nil {
		err = errors.Trace(err)
		return
	}

	ok = len(translations) > 0
	return
}
