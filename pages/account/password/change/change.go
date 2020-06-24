package change

import (
	"github.com/espal-digital-development/espal-core/pages/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/template/renderer"
)

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, form base.Form) Template
}

// Change page service.
type Change struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (change *Change) NewPage(context contexts.Context, form base.Form) Template {
	page := &Page{
		form:            form,
		rendererService: change.rendererService,
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
	form            base.Form
	rendererService renderer.Renderer
}

// Render the page writing to the context.
func (page *Page) Render() {
	base.WritePageTemplate(page.GetCoreContext(), page)
}

// New returns a new instance of Change.
func New(rendererService renderer.Renderer) *Change {
	return &Change{
		rendererService: rendererService,
	}
}
