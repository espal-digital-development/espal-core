package runner

import (
	"fmt"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/entitymutators"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/database/fixtures"
	"github.com/espal-digital-development/espal-core/database/migrations"
	"github.com/juju/errors"
)

const dbURLBlueprint = "postgresql://%s@%s:%d/%s?ssl=true&sslmode=require&sslrootcert=%s&sslkey=%s&sslcert=%s"

func (r *Runner) database() error {
	var err error
	r.databases.creator = database.New()
	err = r.databases.creator.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			r.services.config.DatabaseMigrator(),
			r.services.config.DatabaseHost(),
			r.services.config.DatabasePort(),
			r.services.config.DatabaseName(),
			r.services.config.DatabaseSSLRootCertificateFile(),
			r.services.config.DatabaseMigratorSSLKeyFile(),
			r.services.config.DatabaseMigratorSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	r.databases.deletor = database.New()
	err = r.databases.deletor.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			r.services.config.DatabaseDeletor(),
			r.services.config.DatabaseHost(),
			r.services.config.DatabasePort(),
			r.services.config.DatabaseName(),
			r.services.config.DatabaseSSLRootCertificateFile(),
			r.services.config.DatabaseDeletorSSLKeyFile(),
			r.services.config.DatabaseDeletorSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	r.databases.inserter = database.New()
	err = r.databases.inserter.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			r.services.config.DatabaseInserter(),
			r.services.config.DatabaseHost(),
			r.services.config.DatabasePort(),
			r.services.config.DatabaseName(),
			r.services.config.DatabaseSSLRootCertificateFile(),
			r.services.config.DatabaseInserterSSLKeyFile(),
			r.services.config.DatabaseInserterSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	r.databases.selecter = database.New()
	err = r.databases.selecter.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			r.services.config.DatabaseSelecter(),
			r.services.config.DatabaseHost(),
			r.services.config.DatabasePort(),
			r.services.config.DatabaseName(),
			r.services.config.DatabaseSSLRootCertificateFile(),
			r.services.config.DatabaseSelecterSSLKeyFile(),
			r.services.config.DatabaseSelecterSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	r.databases.updater = database.New()
	err = r.databases.updater.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			r.services.config.DatabaseUpdater(),
			r.services.config.DatabaseHost(),
			r.services.config.DatabasePort(),
			r.services.config.DatabaseName(),
			r.services.config.DatabaseSSLRootCertificateFile(),
			r.services.config.DatabaseUpdaterSSLKeyFile(),
			r.services.config.DatabaseUpdaterSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}

	migratorDatabase := database.New()
	err = migratorDatabase.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			r.services.config.DatabaseMigrator(),
			r.services.config.DatabaseHost(),
			r.services.config.DatabasePort(),
			r.services.config.DatabaseName(),
			r.services.config.DatabaseSSLRootCertificateFile(),
			r.services.config.DatabaseMigratorSSLKeyFile(),
			r.services.config.DatabaseMigratorSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}

	r.services.entityMutators = entitymutators.New(r.databases.inserter, r.databases.updater)
	r.services.databaseFilters = filters.New(r.databases.selecter)
	migrationsService := migrations.New(migratorDatabase, r.services.logger)
	migrationsToRun, err := migrationsService.MigrationsToRun()
	if err != nil {
		return errors.Trace(err)
	}
	if migrationsToRun > 0 {
		r.services.logger.Info("Running migrations..")
	}
	if err := migrationsService.Run(); err != nil {
		return errors.Trace(err)
	}
	fixturesService, err := fixtures.New(r.databases.inserter, r.databases.updater, r.repositories.languages, r.repositories.countries, r.repositories.currencies, r.repositories.userRights)
	if err != nil {
		return errors.Trace(err)
	}
	// Run fixtures only after when it's the first migration run
	if migrationsToRun == migrationsService.TotalMigrations() {
		r.services.logger.Info("Running fixtures..")
		if err := fixturesService.Run(); err != nil {
			return errors.Trace(err)
		}
	}
	// No longer need the migrator
	if err := migratorDatabase.Close(); err != nil {
		return errors.Trace(err)
	}
	return nil
}
