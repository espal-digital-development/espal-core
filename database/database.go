package database

// Database represents an SQL database engine object.
type Database interface {
	Open(driver string, dsn string) error
	Close() error
	Begin() (Transaction, error)
	Exec(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
}
