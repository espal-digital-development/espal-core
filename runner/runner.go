package runner

import (
	"regexp"
	"time"

	"github.com/espal-digital-development/espal-core/cachesynchronizer"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/modules"
	"github.com/espal-digital-development/espal-core/sessions"
	"github.com/juju/errors"
)

var _ Application = &Runner{}

const cacheSynchronizerIntervalInSeconds = 20

// Application represents an object that runs app instances.
type Application interface {
	SetPath(path string)
	Run() error
	RunNonBlocking() error
}

// Runner is an application core that houses and manages all the application's services and their interactions.
type Runner struct {
	path            string
	serverStartTime time.Time

	modulesRegistry []modules.Modular
	reValidSemver   *regexp.Regexp

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

	r.services.cacheSynchronizer = cachesynchronizer.New(r.services.logger, r.stores.cacheNotify,
		time.Second*cacheSynchronizerIntervalInSeconds)

	r.services.sessions = sessions.New(r.services.config, r.stores.session)

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
		modulesRegistry: []modules.Modular{},
		services:        &services{},
		storages:        &storages{},
		databases:       &databases{},
		repositories:    &repositories{},
		stores:          &stores{},
	}
	r.services.logger = logger.New()
	var err error
	r.reValidSemver, err = regexp.Compile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-]` +
		`[0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
