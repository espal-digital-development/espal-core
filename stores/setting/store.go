package setting

import (
	"database/sql"
	"sync"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

// SettingsStore data store.
type SettingsStore struct {
	selecterDatabase database.Database

	// map[siteID][siteID][domainID]themeName
	themesNormal map[string]map[string]map[string]string
	mutex        *sync.RWMutex
}

// GetOneForSite returns the Value for the Site level.
func (s *SettingsStore) GetOneForSite(key uint16, userID string, domainID string, siteID string) (string, error) {
	// TODO :: 77777 Move this caching to the general cache notifier
	if s.mutex == nil {
		s.mutex = &sync.RWMutex{}
		s.themesNormal = make(map[string]map[string]map[string]string)
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if v, ok := s.themesNormal[userID][domainID][siteID]; ok {
		return v, nil
	}

	var err error
	var value string
	if userID == "" {
		value, err = s.getOneForSiteWithoutUser(key, domainID, siteID)
		if err != nil {
			return "", errors.Trace(err)
		}
	} else {
		err = s.selecterDatabase.QueryRow(`SELECT "value" FROM "Setting" WHERE "key" = $1 AND ("userID" = $2
		OR "domainID" = $3 OR "siteID" = $4) ORDER BY "userID", "domainID", "siteID" LIMIT 1`,
			key, userID, domainID, siteID).Scan(&value)
		if err != nil && err != sql.ErrNoRows {
			return "", errors.Trace(err)
		}
	}

	if _, ok := s.themesNormal[userID]; !ok {
		s.themesNormal[userID] = map[string]map[string]string{}
	}
	if _, ok := s.themesNormal[userID][domainID]; !ok {
		s.themesNormal[userID][domainID] = map[string]string{}
	}

	s.themesNormal[userID][domainID][siteID] = value
	return value, nil
}

func (s *SettingsStore) getOneForSiteWithoutUser(key uint16, domainID string, siteID string) (string, error) {
	var value string
	err := s.selecterDatabase.QueryRow(`SELECT "value" FROM "Setting" WHERE "key" = $1 AND ("domainID" = $2
		OR "siteID" = $3) ORDER BY "domainID", "siteID" LIMIT 1`,
		key, domainID, siteID).Scan(&value)
	if err != nil && err != sql.ErrNoRows {
		return "", errors.Trace(err)
	}
	return value, nil
}

// // GetAllForDomainAndSite returns the Values inherited for the Site and Domain level.
// func (s *SettingsStore) GetAllForDomainAndSite(domainID string, siteID string)
//   (settings map[string]string, err error) {
// 	settings = map[string]string{}
// 	rows, err := s.selecterDatabase.Query(`SELECT "key", "value" FROM "Setting"
// 	WHERE ("domainID" = $1 OR "siteID" = $2)
// 	ORDER BY "domainID", "siteID"`, domainID, siteID)
// 	if err == sql.ErrNoRows {
// 		err = nil
// 		return
// 	}
// 	if err != nil {
// 		err = errors.Trace(err)
// 		return
// 	}
// 	defer func(dbRows database.Rows) {
// 		closeErr := dbRows.Close()
// 		if err != nil && closeErr != nil {
// 			err = errors.Wrap(err, closeErr)
// 		} else if closeErr != nil {
// 			err = errors.Trace(closeErr)
// 		}
// 	}(rows)
// 	for rows.Next() {
// 		if err = rows.Err(); err != nil {
// 			err = errors.Trace(err)
// 			return
// 		}
// 		var key, value string
// 		err = rows.Scan(&key, &value)
// 		if err != nil {
// 			err = errors.Trace(err)
// 			return
// 		}
// 		settings[key] = value
// 	}
// 	return
// }
