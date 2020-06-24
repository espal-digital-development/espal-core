package database

import "database/sql"

// Result represents an SQL database result object.
type Result interface {
	sql.Result
}
