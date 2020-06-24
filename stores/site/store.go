package site

import (
	"database/sql"
	"strings"
	"sync"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/juju/errors"
)

type language interface {
	ID() uint16
}

// SitesStore data store.
type SitesStore struct {
	selecterDatabase       database.Database
	updaterDatabase        database.Database
	deletorDatabase        database.Database
	databaseFiltersFactory filters.Factory
	translationsRepository translations.Repository
	loggerService          logger.Loggable
	sitesNormal            map[string]*Site
	mutex                  *sync.RWMutex
}

// GetOne fetches by ID.
func (sitesStore *SitesStore) GetOne(id string) (*Site, bool, error) {
	result, ok, err := sitesStore.fetch(`SELECT * FROM "Site" WHERE "id" = $1 LIMIT 1`, false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneOnlineByID fetches by ID and must be online.
func (sitesStore *SitesStore) GetOneOnlineByID(id string) (*Site, bool, error) {
	// TODO :: 77777 :: Move this caching to the general cache notifier
	if sitesStore.mutex == nil {
		sitesStore.mutex = &sync.RWMutex{}
	}
	sitesStore.mutex.Lock()
	defer sitesStore.mutex.Unlock()

	if sitesStore.sitesNormal == nil {
		sitesStore.sitesNormal = make(map[string]*Site)
	}

	if _, ok := sitesStore.sitesNormal[id]; ok {
		return sitesStore.sitesNormal[id], true, nil
	}

	result, ok, err := sitesStore.fetch(`SELECT * FROM "Site" WHERE "id" = $1 AND "online" = true LIMIT 1`, false, id)
	if len(result) == 1 {
		sitesStore.sitesNormal[id] = result[0]
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// All returns all Sites.
func (sitesStore *SitesStore) All() ([]*Site, bool, error) {
	result, ok, err := sitesStore.fetch(`SELECT * FROM "Site"`, false)
	if err != nil {
		return nil, ok, errors.Trace(err)
	}
	if !ok {
		return nil, ok, nil
	}
	return result, ok, errors.Trace(err)
}

// HasUser returns if Site has User connected to it.
func (sitesStore *SitesStore) HasUser(siteID string, userID string) (bool, error) {
	var id string
	err := sitesStore.selecterDatabase.QueryRow(`SELECT "id" FROM "SiteUser" WHERE "siteID" = $1 AND "userID" = $2 LIMIT 1`, siteID, userID).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Trace(err)
	}
	return id != "", nil
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (sitesStore *SitesStore) GetOneByIDWithCreator(id string) (*Site, bool, error) {
	result, ok, err := sitesStore.fetch(`SELECT s.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "Site" s
		LEFT JOIN "User" cu ON cu."id" = s."createdByID"
		LEFT JOIN "User" uu ON uu."id" = s."updatedByID"
		WHERE s."id" = $1 LIMIT 1`, true, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetTranslatedName returns the presentable name.
func (sitesStore *SitesStore) GetTranslatedName(site *Site, languageID uint16) string {
	var name string
	err := sitesStore.selecterDatabase.QueryRow(`SELECT "value" FROM "SiteTranslation" WHERE "siteID" = $1 AND "field" = $2 AND "language" = $3`, site.ID(), database.DBTranslationFieldName, languageID).Scan(&name)
	if err != nil && err != sql.ErrNoRows {
		sitesStore.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	if err == sql.ErrNoRows || name == "" {
		return sitesStore.translationsRepository.Singular(languageID, "site") + " " + site.ID()
	}
	return name
}

// Delete deletes the given ID(s).
func (sitesStore *SitesStore) Delete(ids []string) error {
	transaction, err := sitesStore.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(`DELETE FROM "SiteTranslation" WHERE "siteID" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(`DELETE FROM "Site" WHERE "id" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// ToggleOnline toggles the active state of the given ID(s).
func (sitesStore *SitesStore) ToggleOnline(ids []string) error {
	transaction, err := sitesStore.updaterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Query(`UPDATE "User" SET "online" = NOT "online" WHERE "id" IN (` + strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}
