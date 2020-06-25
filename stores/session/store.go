package session

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

// SessionsStore data store.
type SessionsStore struct {
	selecterDatabase database.Database
	inserterDatabase database.Database
	updaterDatabase  database.Database
}

// HashExists checks if the record for the given hash exists in the store.
func (s *SessionsStore) HashExists(hash string) (bool, error) {
	row := s.selecterDatabase.QueryRow(`SELECT EXISTS(SELECT 1 FROM "Session" WHERE "hash" = $1 LIMIT 1)`, hash)
	var ok bool
	err := row.Scan(&ok)
	return ok, errors.Trace(err)
}

// GetOne fetches by ID.
func (s *SessionsStore) GetOne(id string) (*Session, bool, error) {
	result, ok, err := s.fetch(`SELECT * FROM "Session" WHERE "id" = $1 LIMIT 1`, false, id)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// GetOneByHash fetches by Hash.
func (s *SessionsStore) GetOneByHash(hash string) (*Session, bool, error) {
	result, ok, err := s.fetch(`SELECT * FROM "Session" WHERE "hash" = $1 LIMIT 1`, false, hash)
	if len(result) == 1 {
		return result[0], ok, errors.Trace(err)
	}
	return nil, ok, errors.Trace(err)
}

// Create creates a new Session instance in the database.
func (s *SessionsStore) Create(hash string, timeout time.Duration, dataEntries DataEntries) error {
	// TODO :: 7777 This is weird/ugly. Need a good uniform style for these interactions
	session := newSession()
	if err := session.SetDataFromJSON(dataEntries); err != nil {
		return errors.Trace(err)
	}
	_, err := s.inserterDatabase.Exec(`INSERT INTO "Session"("hash","timeout","data") VALUES($1,$2,$3)`, hash, timeout, session.Data())
	return errors.Trace(err)
}

// Update updates an existing Session instance in the database.
func (s *SessionsStore) Update(hash string, timeout time.Duration, dataEntries DataEntries) error {
	session := newSession()
	if err := session.SetDataFromJSON(dataEntries); err != nil {
		return errors.Trace(err)
	}
	_, err := s.updaterDatabase.Exec(`UPDATE "Session" SET "timeout" = $1, "data" = $2 WHERE "hash" = $3`, timeout, session.Data(), hash)
	return errors.Trace(err)
}
