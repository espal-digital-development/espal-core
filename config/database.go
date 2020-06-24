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
func (config *Configuration) DatabaseHost() string {
	return config.database.Host
}

// DatabasePort returns the database server port.
func (config *Configuration) DatabasePort() int {
	return config.database.Port
}

// DatabaseName returns the database name.
func (config *Configuration) DatabaseName() string {
	return config.database.Name
}

// DatabaseSelecter returns the database selecter username.
func (config *Configuration) DatabaseSelecter() string {
	return config.database.Users.Selecter
}

// DatabaseCreator returns the database creator username.
func (config *Configuration) DatabaseCreator() string {
	return config.database.Users.Creator
}

// DatabaseInserter returns the database inserter username.
func (config *Configuration) DatabaseInserter() string {
	return config.database.Users.Inserter
}

// DatabaseUpdater returns the database updater username.
func (config *Configuration) DatabaseUpdater() string {
	return config.database.Users.Updater
}

// DatabaseDeletor returns the database deletor username.
func (config *Configuration) DatabaseDeletor() string {
	return config.database.Users.Deletor
}

// DatabaseMigrator returns the database migrator username.
func (config *Configuration) DatabaseMigrator() string {
	return config.database.Users.Migrator
}
