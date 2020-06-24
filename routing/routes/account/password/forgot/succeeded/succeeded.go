package succeeded

import (
	"net/http"

	page "github.com/espal-digital-development/espal-core/pages/account/password/forgot/succeeded"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

// Route processor.
type Route struct {
	succeededPageFactory page.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	if context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	route.succeededPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(succeededPageFactory page.Factory) *Route {
	return &Route{
		succeededPageFactory: succeededPageFactory,
	}
}
