package servererror

import (
	"github.com/espal-digital-development/espal-core/pages/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

var _ Factory = &ServerError{}

// Factory represents an object that serves new pages.
type Factory interface {
	RenderPage(context contexts.Context, title string, message string)
}

// ServerError page service.
type ServerError struct{}

// RenderPage generates a new instance of Page based on the given parameters
// but instantly renders it to prevent package inclusion cycle problems.
func (serverError *ServerError) RenderPage(context contexts.Context, title string, message string) {
	page := &page{
		title:   title,
		message: message,
	}
	page.SetCoreContext(context)
	page.render()
}

type page struct {
	base.Page
	title   string
	message string
}

func (page *page) render() {
	base.WritePageTemplate(page.GetCoreContext(), page)
}

// New returns a new instance of ServerError.
func New() *ServerError {
	return &ServerError{}
}
