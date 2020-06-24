package overview

import (
	page "github.com/espal-digital-development/espal-core/pages/forum/overview"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	forumStore          forum.Store
	overviewPageFactory page.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	forums, _, err := route.forumStore.GetTopLevel(language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	route.overviewPageFactory.NewPage(context, forums, language).Render()
}

// New returns a new instance of Route.
func New(forumStore forum.Store, overviewPageFactory page.Factory) *Route {
	return &Route{
		forumStore:          forumStore,
		overviewPageFactory: overviewPageFactory,
	}
}
