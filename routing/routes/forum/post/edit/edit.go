package edit

import (
	page "github.com/espal-digital-development/espal-core/pages/forum/post/edit"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	regularExpressionsRepository regularexpressions.Repository
	forumStore                   forum.Store
	editPageFactory              page.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	id := context.QueryValue("id")
	if !route.regularExpressionsRepository.GetRouteIDs().MatchString(id) {
		context.RenderBadRequest()
		return
	}
	post, ok, err := route.forumStore.GetOnePostByID(id)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}
	// TODO :: Form, Post and Process
	route.editPageFactory.NewPage(context, post).Render()
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, forumStore forum.Store, editPageFactory page.Factory) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		forumStore:                   forumStore,
		editPageFactory:              editPageFactory,
	}
}
