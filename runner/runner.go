package runner

import (
	"time"

	"github.com/espal-digital-development/espal-core/cachesynchronizer"
	"github.com/espal-digital-development/espal-core/sessions"
	"github.com/juju/errors"
)

var _ Application = &Runner{}

// Application represents an object that runs app instances.
type Application interface {
	Run(path string) error
}

// Runner is an application core that houses and manages all
// the application's services and their interactions.
type Runner struct {
	serverStartTime time.Time

	services       services
	storages       storages
	databases      databases
	repositories   repositories
	stores         stores
	formValidators forms
}

// Run executes all tasks that are needed to startup the application.
func (runner *Runner) Run(path string) error {
	if err := runner.core(path); err != nil {
		return errors.Trace(err)
	}
	if err := runner.database(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.dataStores(); err != nil {
		return errors.Trace(err)
	}

	// TODO :: 777 Make interval a config value
	runner.services.cacheSynchronizer = cachesynchronizer.New(runner.services.logger, runner.stores.cacheNotify, time.Minute)

	runner.services.sessions = sessions.New(runner.services.config, runner.stores.session)

	runner.router()
	runner.forms()
	if err := runner.routes(); err != nil {
		return errors.Trace(err)
	}

	runner.serverStartTime = time.Now()
	runner.startRedirectNonTLSServer()
	runner.startTLSServer()
	runner.listenToSystemSignals()

	return nil
}

// New returns a new instance of Runner.
func New() *Runner {
	return &Runner{}
}
