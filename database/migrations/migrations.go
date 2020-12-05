package migrations

import (
	"fmt"
	"sort"
	"strings"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/migrations/migrationsdata"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/juju/errors"
)

var _ Set = &Migrations{}

// Set represents a collection of database migrations.
type Set interface {
	Run() error
	MigrationsToRun() (uint, error)
	TotalMigrations() uint
}

// Migrations contains all database migrations logic.
type Migrations struct {
	migratorDatabase database.Database
	loggerService    logger.Loggable
}

// Run runs all migrations for the system.
func (m *Migrations) Run() error {
	migrationsToRun, err := m.migrationsToRun()
	if err != nil {
		return errors.Trace(err)
	}

	for _, migration := range migrationsToRun {
		dirName := "_data/" + migration
		files, err := migrationsdata.AssetDir(dirName)
		if err != nil {
			return errors.Trace(err)
		}

		sort.Strings(files)

		var sqlFiles uint16
		for k := range files {
			if files[k][len(files[k])-4:] == ".sql" {
				sqlFiles++
			}
		}

		if sqlFiles == 0 {
			m.loggerService.Infof("Migration `%s` is empty. Skipping..", migration)
			continue
		} else {
			m.loggerService.Infof("Running migration `%s`..", migration)
		}

		var migrationSQL []byte
		for k := range files {
			if files[k][len(files[k])-4:] != ".sql" {
				continue
			}
			migrationSQL, err = migrationsdata.Asset(dirName + "/" + files[k])
			if err != nil {
				return errors.Trace(err)
			}

			m.loggerService.Infof("Running migration `%s` file `%s`..", migration, files[k])

			if _, err := m.migratorDatabase.Exec(string(migrationSQL)); err != nil {
				return errors.Annotate(err, fmt.Sprintf("migration `%s` in file `%s`", migration, files[k]))
			}
		}

		if _, err := m.migratorDatabase.Exec(`INSERT INTO "Migration"("revision") VALUES($1)`, migration); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

// MigrationsToRun returns the amount of migrations that will be run if Run() would be triggered.
func (m *Migrations) MigrationsToRun() (uint, error) {
	migrationsToRun, err := m.migrationsToRun()
	return uint(len(migrationsToRun)), errors.Trace(err)
}

// TotalMigrations returns the total amount of migrations the system has.
func (m *Migrations) TotalMigrations() uint {
	directories := make(map[string]bool)
	allFiles := migrationsdata.AssetNames()
	sort.Strings(allFiles)
	for k := range allFiles {
		file := allFiles[k]
		directory := strings.SplitN(strings.TrimPrefix(file, "_data/"), "/", 2)[0]
		directories[directory] = true
	}
	return uint(len(directories))
}

func (m *Migrations) migrationsToRun() ([]string, error) {
	var lastMigration string
	var migrationsToRun []string
	if err := m.migratorDatabase.QueryRow(`SELECT "revision" FROM "Migration" ORDER BY "id" DESC LIMIT 1`).
		Scan(&lastMigration); err != nil {
		if !strings.Contains(err.Error(), `relation "Migration" does not exist`) {
			return nil, errors.Trace(err)
		}
	}

	if lastMigration == "" {
		migrationsToRun = append(migrationsToRun, "000")
		lastMigration = "000"
	}

	directories := make(map[string]bool)
	allFiles := migrationsdata.AssetNames()
	sort.Strings(allFiles)
	for k := range allFiles {
		file := allFiles[k]
		directory := strings.SplitN(strings.TrimPrefix(file, "_data/"), "/", 2)[0]
		directories[directory] = true
	}

	var targetAfterThis bool
	for directory := range directories {
		if targetAfterThis {
			migrationsToRun = append(migrationsToRun, directory)
			break
		} else if lastMigration == directory {
			targetAfterThis = true
		}
	}

	return migrationsToRun, nil
}

// New returns a new instance of Fixtures.
func New(migratorDatabase database.Database, loggerService logger.Loggable) *Migrations {
	m := &Migrations{
		migratorDatabase: migratorDatabase,
		loggerService:    loggerService,
	}
	return m
}
