package runner

import (
	"time"

	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/modules"
	"github.com/espal-digital-development/espal-core/semver"
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

// Runner is an application core that houses and manages all the application's services and their interactions.
type Runner struct {
	version         string
	path            string
	serverStartTime time.Time

	modulesRegistry []modules.Modular

	services     *services
	storages     *storages
	databases    *databases
	repositories *repositories
	stores       *stores
}

// SetPath sets the app run path.
func (r *Runner) SetPath(path string) {
	r.path = path
}

// Run executes all tasks that are needed to startup the application and blocks code progression.
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

	// TODO :: 777777 Implement when the branch is merged again
	// r.services.notifier = notifier.New(r.services.logger, r.stores.notification)

	// TODO :: 777777 Remove after dev
	// fmt.Println(r.services.notifier.NotifyWithValue(notification.TargetDomain, "blaa", "blup"))

	r.services.sessions = sessions.New(r.services.config, r.stores.session)

	if err := r.themes(); err != nil {
		return errors.Trace(err)
	}

	if err := r.repos(); err != nil {
		return errors.Trace(err)
	}

	r.router()
	if err := r.routes(); err != nil {
		return errors.Trace(err)
	}

	r.serverStartTime = time.Now()
	r.startRedirectNonTLSServer()
	r.startTLSServer()

	r.services.logger.Infof("Server running TLS on `%s` port %d and redirecting non-TLS from port %d",
		r.services.config.ServerHost(), r.services.config.ServerPort(), r.services.config.ServerHTTPRedirectPort())

	r.listenToSystemSignals()

	return nil
}

// New returns a new instance of Runner.
func New() (*Runner, error) {
	r := &Runner{
		version:         "0.0.1",
		modulesRegistry: []modules.Modular{},
		services:        &services{},
		storages:        &storages{},
		databases:       &databases{},
		repositories:    &repositories{},
		stores:          &stores{},
	}
	r.services.logger = logger.New()
	var err error
	r.services.semver, err = semver.New()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
