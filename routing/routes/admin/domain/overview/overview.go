package overview

import (
	"github.com/espal-digital-development/espal-core/pageactions"
	page "github.com/espal-digital-development/espal-core/pages/admin/domain/overview"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/domain"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	domainStore         domain.Store
	overviewPageFactory page.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadDomain") {
		return
	}

	domains, filter, err := route.domainStore.Filter(context)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var canUpdate bool
	var canDelete bool
	if filter.HasResults() {
		canUpdate = context.HasUserRight("UpdateDomain")
		canDelete = context.HasUserRight("DeleteDomain")
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	pageActions := pageactions.New(context, "Domain", len(domains) > 0)
	pageActions.AddCreate()
	pageActions.AddToggle()
	pageActions.AddDelete()
	route.overviewPageFactory.NewPage(context, language, pageActions, filter, domains, canUpdate, canDelete).Render()
}

// New returns a new instance of Route.
func New(domainStore domain.Store, overviewPageFactory page.Factory) *Route {
	return &Route{
		domainStore:         domainStore,
		overviewPageFactory: overviewPageFactory,
	}
}
