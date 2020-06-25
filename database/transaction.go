package database

import (
	"database/sql"
)

// Transaction represents an SQL database transaction object.
type Transaction interface {
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	Rollback() error
	Commit() error
}

type transaction struct {
	tx *sql.Tx
}

// Exec executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
func (t *transaction) Exec(query string, args ...interface{}) (Result, error) {
	return t.tx.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (t *transaction) Query(query string, args ...interface{}) (Rows, error) {
	return t.tx.Query(query, args...)
}

// Rollback aborts the transaction.
func (t *transaction) Rollback() error {
	return t.tx.Rollback()
}

// Commit commits the transaction.
func (t *transaction) Commit() error {
	return t.tx.Commit()
}
