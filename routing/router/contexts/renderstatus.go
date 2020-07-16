package contexts

import (
	"net/http"
	"strconv"

	"github.com/juju/errors"
)

const non200Style = `body{font-family:open sans,helvetica neue,Helvetica,Arial,sans-serif;font-size:13px;` +
	`background:#1c1c1c;color:#eee;}`

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
func (c *HTTPContext) RenderBadRequest() {
	c.SetStatusCode(http.StatusBadRequest)
	c.RenderNon200()
}

// RenderUnauthorized renders a basic 401 page.
func (c *HTTPContext) RenderUnauthorized() {
	c.SetStatusCode(http.StatusUnauthorized)
	c.RenderNon200()
}

// RenderNotFound renders a basic 404 page.
func (c *HTTPContext) RenderNotFound() {
	c.SetStatusCode(http.StatusNotFound)
	c.RenderNon200()
}

// RenderInternalServerError renders a basic 500 page.
func (c *HTTPContext) RenderInternalServerError(err error) {
	c.loggerService.Error(errors.ErrorStack(err))
	c.SetStatusCode(http.StatusInternalServerError)
	var errorMessage string
	if c.configService.Development() {
		errorMessage = err.Error()
	}
	c.RenderNon200Custom(strconv.Itoa(c.StatusCode()), errorMessage)
}

// RenderNon200 gives the possibility to render a non-200 page.
func (c *HTTPContext) RenderNon200() {
	// TODO :: 7 Translate all status messages
	c.RenderNon200Custom(strconv.Itoa(c.StatusCode()), http.StatusText(c.StatusCode()))
}

// RenderNon200Custom gives the possibility to render a non-200 page with a custom message.
// nolint:errcheck
func (c *HTTPContext) RenderNon200Custom(title string, message string) {
	c.WriteString(`<html><head><title>`)
	c.WriteString(title)
	c.WriteString(`</title><style>`)
	c.WriteString(non200Style)
	c.WriteString(`</style></head><body><h1>`)
	c.WriteString(title)
	c.WriteString(`</h1><p>`)
	c.WriteString(message)
	c.WriteString(`</p></body></html>`)
}
