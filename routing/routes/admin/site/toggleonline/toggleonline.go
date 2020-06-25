package toggleonline

import (
	"net/http"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	regularExpressionsRepository regularexpressions.Repository
	siteStore                    site.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("UpdateUser") {
		return
	}
	ids := context.QueryValue("ids")
	if !r.regularExpressionsRepository.GetRouteIDs().MatchString(ids) {
		context.RenderBadRequest()
		return
	}
	if err := r.siteStore.ToggleOnline(strings.Split(ids, ",")); err != nil {
		if err := context.SetFlashErrorMessage(context.Translate("onlineTogglingHasFailed") + ": " + err.Error()); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	} else {
		if err := context.SetFlashSuccessMessage(context.Translate("onlineTogglingWasSuccessful")); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	context.Redirect(context.Referer(), http.StatusTemporaryRedirect)
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, siteStore site.Store) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		siteStore:                    siteStore,
	}
}
