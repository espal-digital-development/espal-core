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
func (transaction *transaction) Exec(query string, args ...interface{}) (Result, error) {
	return transaction.tx.Exec(query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (transaction *transaction) Query(query string, args ...interface{}) (Rows, error) {
	return transaction.tx.Query(query, args...)
}

// Rollback aborts the transaction.
func (transaction *transaction) Rollback() error {
	return transaction.tx.Rollback()
}

// Commit commits the transaction.
func (transaction *transaction) Commit() error {
	return transaction.tx.Commit()
}
