package database

import (
	"database/sql"
)

// Rows represents an SQL database rows object.
type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Err() error
	Close() error
}

type rows struct {
	rows *sql.Rows
}

// Next prepares the next result row for reading with the Scan method. It
// returns true on success, or false if there is no next result row or an error
// happened while preparing it. Err should be consulted to distinguish between
// the two cases.
//
// Every call to Scan, even the first one, must be preceded by a call to Next.
func (rows *rows) Next() bool {
	return rows.rows.Next()
}

// Scan copies the columns in the current row into the values pointed
// at by dest. The number of values in dest must be the same as the
// number of columns in Rows.
func (rows *rows) Scan(dest ...interface{}) error {
	return rows.rows.Scan(dest...)
}

// Err returns the error, if any, that was encountered during iteration.
// Err may be called after an explicit or implicit Close.
func (rows *rows) Err() error {
	return rows.rows.Err()
}

// Close closes the Rows, preventing further enumeration. If Next is called
// and returns false and there are no further result sets,
// the Rows are closed automatically and it will suffice to check the
// result of Err. Close is idempotent and does not affect the result of Err.
func (rows *rows) Close() error {
	return rows.rows.Close()
}
