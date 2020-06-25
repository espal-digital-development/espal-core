package setting

import (
	"database/sql"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

// SettingsStore data store.
type SettingsStore struct {
	selecterDatabase database.Database
}

// GetOneForSite returns the Value for the Site level.
func (s *SettingsStore) GetOneForSite(key uint16, userID string, domainID string, siteID string) (string, error) {
	var value string
	err := s.selecterDatabase.QueryRow(`SELECT "value" FROM "Setting" WHERE "key" = $1 AND ("userID" = $2 OR "domainID" = $3 OR "siteID" = $4) ORDER BY "userID", "domainID", "siteID" LIMIT 1`, key, userID, domainID, siteID).Scan(&value)
	if err != nil && err != sql.ErrNoRows {
		return "", errors.Trace(err)
	}
	return value, nil
}
