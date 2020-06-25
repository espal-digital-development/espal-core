package recovery

import (
	"github.com/espal-digital-development/espal-core/pages/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

var _ Factory = &Recovery{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, form base.Form) Template
}

// Recovery page service.
type Recovery struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (r *Recovery) NewPage(context contexts.Context, form base.Form) Template {
	page := &Page{
		form: form,
	}
	page.SetCoreContext(context)
	return page
}

// Template represents a renderable page template object.
type Template interface {
	Render()
}

// Page contains and handles template logic.
type Page struct {
	base.Page
	form base.Form
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Recovery.
func New() *Recovery {
	return &Recovery{}
}
