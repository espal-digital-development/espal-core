package auth

import (
	"net/http"

	page "github.com/espal-digital-development/espal-core/pages/auth"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators/forms/auth"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	authFormValidator auth.Factory
	authPageFactory   page.Factory
}

// Handle route handler.
func (route *Route) Handle(context contexts.Context) {
	if context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	// TODO :: 7 This is ok for now, but the Auth should be blocking the site to such an extent that it shouldn't even show the design or expose the assets.
	form, err := route.authFormValidator.New(language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	defer form.Close()
	isSubmitted, isValid, err := form.Submit(context)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if isSubmitted && isValid {
		err = context.Login(form.GetUserID(), form.RememberMe())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	route.authPageFactory.NewPage(context, form.View()).Render()
}

// New returns a new instance of Route.
func New(authFormValidator auth.Factory, authPageFactory page.Factory) *Route {
	return &Route{
		authFormValidator: authFormValidator,
		authPageFactory:   authPageFactory,
	}
}
