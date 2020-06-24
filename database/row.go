package database

import (
	"database/sql"
)

// Row represents an SQL database row object.
type Row interface {
	Scan(dest ...interface{}) error
}

type row struct {
	row *sql.Row
}

// Scan copies the columns from the matched row into the values
// pointed at by dest. See the documentation on Rows.Scan for details.
// If more than one row matches the query,
// Scan uses the first row and discards the rest. If no row matches
// the query, Scan returns ErrNoRows.
func (row *row) Scan(dest ...interface{}) error {
	return row.row.Scan(dest...)
}
