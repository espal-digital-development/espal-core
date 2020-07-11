package domain

import (
	"strings"
	"sync"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/database/queryhelper"
	"github.com/juju/errors"
)

// DomainsStore data store.
type DomainsStore struct {
	selecterDatabase       database.Database
	updaterDatabase        database.Database
	deletorDatabase        database.Database
	databaseQueryHelper    queryhelper.Helper
	databaseFiltersFactory filters.Factory
	domainsNormal          map[string]*Domain
	mutex                  *sync.RWMutex
}

// GetOne fetches by ID.
func (d *DomainsStore) GetOne(id string) (*Domain, bool, error) {
	result, ok, err := d.fetch(`SELECT * FROM "Domain" WHERE "id" = $1 LIMIT 1`, false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// All returns all Domains.
func (d *DomainsStore) All() ([]*Domain, bool, error) {
	result, ok, err := d.fetch(`SELECT * FROM "Domain"`, false)
	if err != nil {
		return nil, ok, errors.Trace(err)
	}
	if !ok {
		return nil, ok, nil
	}
	return result, ok, errors.Trace(err)
}

// GetOneByIDWithCreator fetches by ID, including the CreatedBy and UpdatedBy fields.
func (d *DomainsStore) GetOneByIDWithCreator(id string) (*Domain, bool, error) {
	result, ok, err := d.fetch(`SELECT d.*, cu."firstName", cu."surname", uu."firstName", uu."surname"
		FROM "Domain" d
		LEFT JOIN "User" cu ON cu."id" = d."createdByID"
		LEFT JOIN "User" uu ON uu."id" = d."updatedByID"
		WHERE d."id" = $1 LIMIT 1`, true, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneActiveByHost fetches by Host and must be Active.
func (d *DomainsStore) GetOneActiveByHost(host string) (*Domain, bool, error) {
	// TODO :: 77777 Move this caching to the general cache notifier
	if d.mutex == nil {
		d.mutex = &sync.RWMutex{}
		d.domainsNormal = make(map[string]*Domain)
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if _, ok := d.domainsNormal[host]; ok {
		return d.domainsNormal[host], true, nil
	}

	result, ok, err := d.fetch(`SELECT * FROM "Domain" WHERE "host" = $1 AND "active" = true LIMIT 1`, false, host)
	if len(result) == 1 {
		d.domainsNormal[host] = result[0]
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// Delete deletes the given ID(s).
func (d *DomainsStore) Delete(ids []string) error {
	transaction, err := d.deletorDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	query, idsInterfaces, err := d.databaseQueryHelper.BuildDeleteWhereInIds("Domain", "id", ids)
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

// ToggleActive toggles the active state of the given ID(s).
func (d *DomainsStore) ToggleActive(ids []string) error {
	transaction, err := d.updaterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	if _, err := transaction.Query(`UPDATE "Domain" SET "active" = NOT "active" WHERE "id" IN (` +
		strings.Join(ids, ",") + `)`); err != nil {
		if err := transaction.Rollback(); err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	}
	return transaction.Commit()
}
