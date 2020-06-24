package contexts

import (
	"net/http"
	"strconv"

	"github.com/juju/errors"
)

// RenderStatusContext for status rendering.
type RenderStatusContext interface {
	RenderBadRequest()
	RenderUnauthorized()
	RenderNotFound()
	RenderInternalServerError(error)
	RenderNon200()
	RenderNon200Custom(title string, message string)
}

// RenderBadRequest renders a basic 400 page.
func (httpContext *HTTPContext) RenderBadRequest() {
	httpContext.SetStatusCode(http.StatusBadRequest)
	httpContext.RenderNon200()
}

// RenderUnauthorized renders a basic 401 page.
func (httpContext *HTTPContext) RenderUnauthorized() {
	httpContext.SetStatusCode(http.StatusUnauthorized)
	httpContext.RenderNon200()
}

// RenderNotFound renders a basic 404 page.
func (httpContext *HTTPContext) RenderNotFound() {
	httpContext.SetStatusCode(http.StatusNotFound)
	httpContext.RenderNon200()
}

// RenderInternalServerError renders a basic 500 page.
func (httpContext *HTTPContext) RenderInternalServerError(err error) {
	httpContext.loggerService.Error(errors.ErrorStack(err))
	httpContext.SetStatusCode(http.StatusInternalServerError)
	var errorMessage string
	if httpContext.configService.Development() {
		errorMessage = err.Error()
	}
	httpContext.RenderNon200Custom(strconv.Itoa(httpContext.StatusCode()), errorMessage)
}

// RenderNon200 gives the possibility to render a non-200 page.
func (httpContext *HTTPContext) RenderNon200() {
	// TODO :: 7 Translate all status messages
	httpContext.RenderNon200Custom(strconv.Itoa(httpContext.StatusCode()), http.StatusText(httpContext.StatusCode()))
}

// RenderNon200Custom gives the possibility to render a non-200 page with a custom message.
func (httpContext *HTTPContext) RenderNon200Custom(title string, message string) {
	httpContext.serverError.RenderPage(httpContext, title, message)
}
