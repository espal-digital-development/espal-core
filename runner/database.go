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

func (runner *Runner) database() error {
	var err error
	runner.databases.creator = database.New()
	err = runner.databases.creator.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			runner.services.config.DatabaseMigrator(),
			runner.services.config.DatabaseHost(),
			runner.services.config.DatabasePort(),
			runner.services.config.DatabaseName(),
			runner.services.config.DatabaseSSLRootCertificateFile(),
			runner.services.config.DatabaseMigratorSSLKeyFile(),
			runner.services.config.DatabaseMigratorSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	runner.databases.deletor = database.New()
	err = runner.databases.deletor.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			runner.services.config.DatabaseDeletor(),
			runner.services.config.DatabaseHost(),
			runner.services.config.DatabasePort(),
			runner.services.config.DatabaseName(),
			runner.services.config.DatabaseSSLRootCertificateFile(),
			runner.services.config.DatabaseDeletorSSLKeyFile(),
			runner.services.config.DatabaseDeletorSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	runner.databases.inserter = database.New()
	err = runner.databases.inserter.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			runner.services.config.DatabaseInserter(),
			runner.services.config.DatabaseHost(),
			runner.services.config.DatabasePort(),
			runner.services.config.DatabaseName(),
			runner.services.config.DatabaseSSLRootCertificateFile(),
			runner.services.config.DatabaseInserterSSLKeyFile(),
			runner.services.config.DatabaseInserterSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	runner.databases.selecter = database.New()
	err = runner.databases.selecter.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			runner.services.config.DatabaseSelecter(),
			runner.services.config.DatabaseHost(),
			runner.services.config.DatabasePort(),
			runner.services.config.DatabaseName(),
			runner.services.config.DatabaseSSLRootCertificateFile(),
			runner.services.config.DatabaseSelecterSSLKeyFile(),
			runner.services.config.DatabaseSelecterSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}
	runner.databases.updater = database.New()
	err = runner.databases.updater.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			runner.services.config.DatabaseUpdater(),
			runner.services.config.DatabaseHost(),
			runner.services.config.DatabasePort(),
			runner.services.config.DatabaseName(),
			runner.services.config.DatabaseSSLRootCertificateFile(),
			runner.services.config.DatabaseUpdaterSSLKeyFile(),
			runner.services.config.DatabaseUpdaterSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}

	migratorDatabase := database.New()
	err = migratorDatabase.Open("postgres",
		fmt.Sprintf(dbURLBlueprint,
			runner.services.config.DatabaseMigrator(),
			runner.services.config.DatabaseHost(),
			runner.services.config.DatabasePort(),
			runner.services.config.DatabaseName(),
			runner.services.config.DatabaseSSLRootCertificateFile(),
			runner.services.config.DatabaseMigratorSSLKeyFile(),
			runner.services.config.DatabaseMigratorSSLCertificateFile(),
		),
	)
	if err != nil {
		return errors.Trace(err)
	}

	runner.services.entityMutators = entitymutators.New(runner.databases.inserter, runner.databases.updater)
	runner.services.databaseFilters = filters.New(runner.databases.selecter)
	migrationsService := migrations.New(migratorDatabase, runner.services.logger)
	migrationsToRun, err := migrationsService.MigrationsToRun()
	if err != nil {
		return errors.Trace(err)
	}
	if migrationsToRun > 0 {
		runner.services.logger.Info("Running migrations..")
	}
	if err := migrationsService.Run(); err != nil {
		return errors.Trace(err)
	}
	fixturesService, err := fixtures.New(runner.databases.inserter, runner.databases.updater, runner.repositories.languages, runner.repositories.countries, runner.repositories.currencies, runner.repositories.userRights)
	if err != nil {
		return errors.Trace(err)
	}
	// Run fixtures only after when it's the first migration run
	if migrationsToRun == migrationsService.TotalMigrations() {
		runner.services.logger.Info("Running fixtures..")
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
