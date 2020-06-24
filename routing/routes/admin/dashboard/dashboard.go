package dashboard

import (
	"github.com/espal-digital-development/espal-core/pages/admin/dashboard"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

// Route processor.
type Route struct {
	dashboardPageFactory dashboard.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	if !context.HasAdminAccess() {
		context.RenderUnauthorized()
		return
	}

	route.dashboardPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(dashboardPageFactory dashboard.Factory) *Route {
	return &Route{
		dashboardPageFactory: dashboardPageFactory,
	}
}
