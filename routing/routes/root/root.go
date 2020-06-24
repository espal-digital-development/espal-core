package root

import (
	"github.com/espal-digital-development/espal-core/pages/root"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

// Route processor.
type Route struct {
	rootPageFactory root.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	route.rootPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(rootPageFactory root.Factory) *Route {
	return &Route{
		rootPageFactory: rootPageFactory,
	}
}
