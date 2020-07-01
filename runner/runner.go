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
	SetPath(path string)
	Run() error
	RunNonBlocking() error
}

// Runner is an application core that houses and manages all
// the application's services and their interactions.
type Runner struct {
	path            string
	serverStartTime time.Time

	services       services
	storages       storages
	databases      databases
	repositories   repositories
	stores         stores
	formValidators forms
}

// SetPath sets the app run path.
func (r *Runner) SetPath(path string) {
	r.path = path
}

// Run executes all tasks that are needed to startup the application
// and blocks code progression.
func (r *Runner) Run() error {
	if err := r.RunNonBlocking(); err != nil {
		return errors.Trace(err)
	}
	select {}
}

// RunNonBlocking executes all tasks that are needed to startup the application.
func (r *Runner) RunNonBlocking() error {
	if r.path == "" {
		r.path = "./app"
	}

	if err := r.core(r.path); err != nil {
		return errors.Trace(err)
	}
	if err := r.database(); err != nil {
		return errors.Trace(err)
	}
	if err := r.dataStores(); err != nil {
		return errors.Trace(err)
	}

	// TODO :: 777 Make interval a config value
	r.services.cacheSynchronizer = cachesynchronizer.New(r.services.logger, r.stores.cacheNotify, time.Minute)

	r.services.sessions = sessions.New(r.services.config, r.stores.session)

	r.router()
	r.forms()
	if err := r.routes(); err != nil {
		return errors.Trace(err)
	}

	r.serverStartTime = time.Now()
	r.startRedirectNonTLSServer()
	r.startTLSServer()
	r.listenToSystemSignals()

	return nil
}

// New returns a new instance of Runner.
func New() *Runner {
	return &Runner{}
}
