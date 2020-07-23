package router

import (
	"net/http"
	"sync"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/domain"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-core/stores/slug"
	"github.com/juju/errors"
)

var _ Router = &HTTPRouter{}

type handler interface {
	Handle(contexts.Context)
}

// Router represents an object that controls routes registration and serving them on an endpoint.
type Router interface {
	RegisterRoute(path string, h handler) error
	UnregisterRoute(path string) error
	ServeHTTP(responseWriter http.ResponseWriter, request *http.Request)
}

// HTTPRouter HTTP routing engine service.
type HTTPRouter struct {
	configService   config.Config
	loggerService   logger.Loggable
	contextsFactory contexts.Factory
	domainStore     domain.Store
	siteStore       site.Store
	slugStore       slug.Store

	routes      map[string]handler
	routesMutex *sync.RWMutex
}

// RegisterRoute registers a callback to a new route.
func (r *HTTPRouter) RegisterRoute(path string, h handler) error {
	r.routesMutex.Lock()
	defer r.routesMutex.Unlock()
	if _, ok := r.routes[path]; ok {
		return errors.Errorf("path `%s` is already registered", path)
	}
	r.routes[path] = h
	return nil
}

// UnregisterRoute unregisters a route and it's callback for the given domain.
func (r *HTTPRouter) UnregisterRoute(path string) error {
	r.routesMutex.Lock()
	defer r.routesMutex.Unlock()
	if _, ok := r.routes[path]; !ok {
		return errors.Errorf("`%s` is not known route", path)
	}
	delete(r.routes, path)
	return nil
}

func (r *HTTPRouter) getRoute(path string) (handler, bool) {
	r.routesMutex.RLock()
	defer r.routesMutex.RUnlock()
	route, ok := r.routes[path]
	return route, ok
}

// New returns a new instance of a Router.
func New(configService config.Config, loggerService logger.Loggable, contextsFactory contexts.Factory, domainStore domain.Store,
	siteStore site.Store, slugStore slug.Store) *HTTPRouter {
	return &HTTPRouter{
		configService:   configService,
		loggerService:   loggerService,
		contextsFactory: contextsFactory,
		domainStore:     domainStore,
		siteStore:       siteStore,
		slugStore:       slugStore,

		routes:      map[string]handler{},
		routesMutex: &sync.RWMutex{},
	}
}
