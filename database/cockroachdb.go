package database

import (
	"database/sql"

	"github.com/juju/errors"
)

var _ Database = &CockroachDB{}

// TODO :: 7777 Eventually needs to move to a `engines/cockroachdb` subpackage but keep it here for now.

// CockroachDB offers driver interaction with a CockroachDB cluster.
type CockroachDB struct {
	db *sql.DB
}

// Open opens a database specified by its database driver name and a driver-specific data source name, usually
// consisting of at least a database name and connection information.
func (c *CockroachDB) Open(driver string, dsn string) error {
	var err error
	c.db, err = sql.Open(driver, dsn)
	return errors.Trace(err)
}

// Close closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
// It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
func (c *CockroachDB) Close() error {
	return c.db.Close()
}

// Begin starts a transaction.
func (c *CockroachDB) Begin() (Transaction, error) {
	var err error
	transaction := &transaction{}
	transaction.tx, err = c.db.Begin()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return transaction, nil
}

// Exec executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
func (c *CockroachDB) Exec(query string, args ...interface{}) (Result, error) {
	return c.db.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (c *CockroachDB) Query(query string, args ...interface{}) (Rows, error) {
	var err error
	dbRows := &rows{}
	// nolint:rowserrcheck
	dbRows.rows, err = c.db.Query(query, args...)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return dbRows, nil
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards the rest.
func (c *CockroachDB) QueryRow(query string, args ...interface{}) Row {
	dbRow := &row{}
	dbRow.row = c.db.QueryRow(query, args...)
	return dbRow
}

// New returns a new of CockroachDB.
func New() *CockroachDB {
	return &CockroachDB{}
}
