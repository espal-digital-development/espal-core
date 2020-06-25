package catalog

import (
	page "github.com/espal-digital-development/espal-core/pages/catalog"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

// Route processor.
type Route struct {
	catalogPageFactory page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	r.catalogPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(catalogPageFactory page.Factory) *Route {
	return &Route{
		catalogPageFactory: catalogPageFactory,
	}
}
