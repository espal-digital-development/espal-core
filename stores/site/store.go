package site

import (
	"database/sql"
	errorsNative "errors"
	"sync"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/database/queryhelper"
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
	databaseQueryHelper    queryhelper.Helper
	databaseFiltersFactory filters.Factory
	translationsRepository translations.Repository
	loggerService          logger.Loggable

	cacheNormal map[string]*Site
	mutex       *sync.RWMutex
}

// GetOne fetches by ID.
func (s *SitesStore) GetOne(id string) (*Site, bool, error) {
	result, ok, err := s.fetch(`SELECT * FROM "Site" WHERE "id" = $1 LIMIT 1`, false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneOnlineByID fetches by ID and must be online.
func (s *SitesStore) GetOneOnlineByID(id string) (*Site, bool, error) {
	// TODO :: 77777 Move this caching to the general cache notifier
	if s.mutex == nil {
		s.mutex = &sync.RWMutex{}
		s.cacheNormal = make(map[string]*Site)
	}
	s.mutex.RLock()

	if v, ok := s.cacheNormal[id]; ok {
		s.mutex.RUnlock()
		return v, true, nil
	}
	s.mutex.RUnlock()

	result, ok, err := s.fetch(`SELECT * FROM "Site" WHERE "id" = $1 AND "online" = true LIMIT 1`, false, id)
	if len(result) == 1 {
		s.mutex.Lock()
		s.cacheNormal[id] = result[0]
		s.mutex.Unlock()
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// All returns all Sites.
func (s *SitesStore) All() ([]*Site, bool, error) {
	result, ok, err := s.fetch(`SELECT * FROM "Site"`, false)
	if err != nil {
		return nil, ok, errors.Trace(err)
	}
	if !ok {
		return nil, ok, nil
	}
	return result, ok, errors.Trace(err)
}

// HasUser returns if Site has User connected to it.
func (s *SitesStore) HasUser(siteID string, userID string) (bool, error) {
	var id string
	err := s.selecterDatabase.QueryRow(`SELECT "id" FROM "SiteUser" WHERE "siteID" = $1 AND "userID" = $2 LIMIT 1`,
		siteID, userID).Scan(&id)
	if err != nil && !errorsNative.Is(err, sql.ErrNoRows) {
		return false, errors.Trace(err)
	}
	return id != "", nil
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (s *SitesStore) GetOneByIDWithCreator(id string) (*Site, bool, error) {
	result, ok, err := s.fetch(`SELECT s.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
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
func (s *SitesStore) GetTranslatedName(site *Site, languageID uint16) string {
	var name string
	err := s.selecterDatabase.QueryRow(`SELECT "value" FROM "SiteTranslation" WHERE "siteID" = $1 AND "field" = $2
		AND "language" = $3`, site.ID(), database.DBTranslationFieldName, languageID).Scan(&name)
	if err != nil && !errorsNative.Is(err, sql.ErrNoRows) {
		s.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	if name == "" || !errorsNative.Is(err, sql.ErrNoRows) {
		return s.translationsRepository.Singular(languageID, "site") + " " + site.ID()
	}
	return name
}

// Delete deletes the given ID(s).
func (s *SitesStore) Delete(ids []string) error {
	transaction, err := s.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	query, idsInterfaces, err := s.databaseQueryHelper.BuildDeleteWhereInIds("SiteTranslation", "siteID", ids)
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(query, idsInterfaces...); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	query, idsInterfaces, err = s.databaseQueryHelper.BuildDeleteWhereInIds("Site", "id", ids)
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(query, idsInterfaces...); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}

// ToggleOnline toggles the active state of the given ID(s).
func (s *SitesStore) ToggleOnline(ids []string) error {
	transaction, err := s.updaterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	query, idsInterfaces, err := s.databaseQueryHelper.BuildUpdateWhereInIds("SiteTranslation",
		`SET "online" = NOT "online"`, "siteID", ids)
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Exec(query, idsInterfaces...); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}
