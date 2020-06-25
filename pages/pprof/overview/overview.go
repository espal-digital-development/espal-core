package overview

import (
	"runtime/pprof"

	"github.com/espal-digital-development/espal-core/pages/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

var _ Factory = &Overview{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, profiles []*pprof.Profile) Template
}

// Overview page service.
type Overview struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (o *Overview) NewPage(context contexts.Context, profiles []*pprof.Profile) Template {
	page := &Page{
		profiles: profiles,
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
	profiles []*pprof.Profile
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Overview.
func New() *Overview {
	return &Overview{}
}
