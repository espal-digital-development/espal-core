package notification

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

// NotificationsStore data store.
type NotificationsStore struct {
	selecterDatabase database.Database
	inserterDatabase database.Database
	deletorDatabase  database.Database
}

// GetLatest fetches the latest entries since the given interval.
func (s *NotificationsStore) GetLatest(interval time.Duration) ([]*Notification, bool, error) {
	result, ok, err := s.fetch(`SELECT DISTINCT ON("target","key") * FROM "Notification" WHERE "createdAt" > NOW()
		- INTERVAL '$1 SECONDS' ORDER BY "target","key","createdAt" DESC`, false, interval.Seconds())
	return result, ok, errors.Trace(err)
}

// Save saves a new entry for the given target, key and value (optional).
func (s *NotificationsStore) Save(target string, key string, value string) (string, error) {
	var insertedID string
	row := s.inserterDatabase.QueryRow(`INSERT INTO "Notification"("target","key","value") VALUES($1,$2,$3)`,
		target, key, value)
	if err := row.Scan(&insertedID); err != nil {
		return "", errors.Trace(err)
	}
	return insertedID, nil
}
