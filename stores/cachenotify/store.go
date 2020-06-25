package cachenotify

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

// CacheNotifiesStore data store.
type CacheNotifiesStore struct {
	selecterDatabase database.Database
	updaterDatabase  database.Database
}

// GetLatest fetches the latest entries since the given interval.
func (c *CacheNotifiesStore) GetLatest(interval time.Duration) ([]*CacheNotify, bool, error) {
	result, ok, err := c.fetch(`SELECT DISTINCT ON("target", "key") * FROM "CacheNotify" WHERE "createdAt" > NOW() - INTERVAL '$1 SECONDS' ORDER BY "target", "key", "createdAt" DESC`, false, interval.Seconds())
	return result, ok, errors.Trace(err)
}

// Save saves a new entry for the given target and key.
func (c *CacheNotifiesStore) Save(target uint, key string) error {
	// TODO :: 7777 Implement
	return nil
}
