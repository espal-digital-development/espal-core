package overview

import (
	"github.com/espal-digital-development/espal-core/pageactions"
	page "github.com/espal-digital-development/espal-core/pages/admin/user/group/overview"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user/group"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	userGroupStore      group.Store
	overviewPageFactory page.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadUserGroup") {
		return
	}
	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	userGroups, filter, err := route.userGroupStore.Filter(context, language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var canUpdate bool
	var canDelete bool
	if filter.HasResults() {
		canUpdate = context.HasUserRight("UpdateUserGroup")
		canDelete = context.HasUserRight("DeleteUserGroup")
	}
	pageActions := pageactions.New(context, "UserGroup", len(userGroups) > 0)
	pageActions.AddCreate()
	pageActions.AddToggle()
	pageActions.AddDelete()
	route.overviewPageFactory.NewPage(context, language, pageActions, filter, userGroups, canUpdate, canDelete).Render()
}

// New returns a new instance of Route.
func New(userGroupStore group.Store, overviewPageFactory page.Factory) *Route {
	return &Route{
		userGroupStore:      userGroupStore,
		overviewPageFactory: overviewPageFactory,
	}
}
