package databasemigrations

// DatabaseMigrations provider object.
type DatabaseMigrations struct{}

// New returns a new instance of DatabaseMigrations.
func New() (*DatabaseMigrations, error) {
	m := &DatabaseMigrations{}
	return m, nil
}
