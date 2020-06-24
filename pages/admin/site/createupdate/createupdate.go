package createupdate

import (
	"github.com/espal-digital-development/espal-core/pages/admin/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-core/template/renderer"
)

var _ Factory = &CreateUpdate{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, site *site.Site, language contexts.Language, form base.Form, displayTitle string) Template
}

// CreateUpdate page service.
type CreateUpdate struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (createUpdate *CreateUpdate) NewPage(context contexts.Context, site *site.Site, language contexts.Language, form base.Form, displayTitle string) Template {
	page := &Page{
		site:            site,
		language:        language,
		form:            form,
		displayTitle:    displayTitle,
		rendererService: createUpdate.rendererService,
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
	site            *site.Site
	language        contexts.Language
	form            base.Form
	displayTitle    string
	rendererService renderer.Renderer
}

// Render the page writing to the context.
func (page *Page) Render() {
	base.WritePageTemplate(page.GetCoreContext(), page)
}

// New returns a new instance of CreateUpdate.
func New(rendererService renderer.Renderer) *CreateUpdate {
	return &CreateUpdate{
		rendererService: rendererService,
	}
}
