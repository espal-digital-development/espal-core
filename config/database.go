package config

// Database config section.
type Database interface {
	DatabaseHost() string
	DatabasePort() int
	DatabaseName() string
	DatabaseSelecter() string
	DatabaseCreator() string
	DatabaseInserter() string
	DatabaseUpdater() string
	DatabaseDeletor() string
	DatabaseMigrator() string
}

type database struct {
	Host  string
	Port  int
	Name  string
	Users struct {
		Selecter string
		Creator  string
		Inserter string
		Updater  string
		Deletor  string
		Migrator string
	}
}

// DatabaseHost returns the database host server address.
func (c *Configuration) DatabaseHost() string {
	return c.database.Host
}

// DatabasePort returns the database server port.
func (c *Configuration) DatabasePort() int {
	return c.database.Port
}

// DatabaseName returns the database name.
func (c *Configuration) DatabaseName() string {
	return c.database.Name
}

// DatabaseSelecter returns the database selecter username.
func (c *Configuration) DatabaseSelecter() string {
	return c.database.Users.Selecter
}

// DatabaseCreator returns the database creator username.
func (c *Configuration) DatabaseCreator() string {
	return c.database.Users.Creator
}

// DatabaseInserter returns the database inserter username.
func (c *Configuration) DatabaseInserter() string {
	return c.database.Users.Inserter
}

// DatabaseUpdater returns the database updater username.
func (c *Configuration) DatabaseUpdater() string {
	return c.database.Users.Updater
}

// DatabaseDeletor returns the database deletor username.
func (c *Configuration) DatabaseDeletor() string {
	return c.database.Users.Deletor
}

// DatabaseMigrator returns the database migrator username.
func (c *Configuration) DatabaseMigrator() string {
	return c.database.Users.Migrator
}
