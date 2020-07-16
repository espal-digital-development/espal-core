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
	// TODO :: 7 This is ok for now, but the Auth should be blocking the site to such an extent
	// that it shouldn't even show the design or expose the assets.
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

	// TODO :: 777777 Implement nice native rendered output
	// 	<div class="simpleBox">
	//     {%s= p.form.Errors() %}
	//     <h1>{%s p.Translate("authentication") %}</h1>
	//     {%s= p.form.Open() %}
	//         {%s= p.form.Field("_uname") %}
	//         {%s= p.form.Field("_t") %}
	//         {%s= p.form.Field("email") %}<br>
	//         {%s= p.form.Field("password") %}<br>
	//         {%s= p.form.Field("rememberMe") %}<br>
	//         <input type="submit" value="{%s p.Translate("login") %}">
	//     </form>
	// </div>
}

// New returns a new instance of Route.
func New(authFormValidator auth.Factory) *Route {
	return &Route{
		authFormValidator: authFormValidator,
	}
}
