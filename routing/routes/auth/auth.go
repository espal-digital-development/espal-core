package auth

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators/forms/auth"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	authFormValidator auth.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	form, err := r.authFormValidator.New(language)
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

	context.WriteString(`<div class="simpleBox">`)
	context.WriteString(form.View().Errors())
	context.WriteString(`<h1>`)
	context.WriteString(context.Translate("authentication"))
	context.WriteString(`</h1>`)
	context.WriteString(form.View().Open())
	context.WriteString(form.View().Field("_uname"))
	context.WriteString(form.View().Field("_t"))
	context.WriteString(form.View().Field("email"))
	context.WriteString(`<br>`)
	context.WriteString(form.View().Field("password"))
	context.WriteString(`<br>`)
	context.WriteString(form.View().Field("rememberMe"))
	context.WriteString(`<br><input type="submit" value="`)
	context.WriteString(context.Translate("login"))
	context.WriteString(`"></form></div>`)
}

// New returns a new instance of Route.
func New(authFormValidator auth.Factory) *Route {
	return &Route{
		authFormValidator: authFormValidator,
	}
}
