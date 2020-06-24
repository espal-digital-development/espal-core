package overview

import (
	"runtime/pprof"

	page "github.com/espal-digital-development/espal-core/pages/pprof/overview"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

// Route processor.
type Route struct {
	overviewPageFactory page.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	route.overviewPageFactory.NewPage(context, pprof.Profiles()).Render()
}

// New returns a new instance of Route.
func New(overviewPageFactory page.Factory) *Route {
	return &Route{
		overviewPageFactory: overviewPageFactory,
	}
}
