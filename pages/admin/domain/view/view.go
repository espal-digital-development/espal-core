package view

import (
	"github.com/espal-digital-development/espal-core/pages/admin/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/domain"
)

var _ Factory = &View{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, language contexts.Language, domain *domain.Domain, domainLanguage contexts.Language) Template
}

// View page service.
type View struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (view *View) NewPage(context contexts.Context, language contexts.Language, domain *domain.Domain, domainLanguage contexts.Language) Template {
	page := &Page{
		language:       language,
		domain:         domain,
		domainLanguage: domainLanguage,
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
	language       contexts.Language
	domain         *domain.Domain
	domainLanguage contexts.Language
}

// Render the page writing to the context.
func (page *Page) Render() {
	base.WritePageTemplate(page.GetCoreContext(), page)
}

// New returns a new instance of View.
func New() *View {
	return &View{}
}
