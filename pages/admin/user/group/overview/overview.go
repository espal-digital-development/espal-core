package overview

import (
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/pages/admin/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user/group"
	"github.com/espal-digital-development/espal-core/template/renderer"
)

var _ Factory = &Overview{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, language contexts.Language, actions pageactions.Actions, filter filters.Filter,
		userGroups []*group.Group, canUpdate bool, canDelete bool) Template
}

// Overview page service.
type Overview struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (o *Overview) NewPage(context contexts.Context, language contexts.Language, actions pageactions.Actions,
	filter filters.Filter, userGroups []*group.Group, canUpdate bool, canDelete bool) Template {
	page := &Page{
		language:        language,
		actions:         actions,
		filter:          filter,
		userGroups:      userGroups,
		canUpdate:       canUpdate,
		canDelete:       canDelete,
		rendererService: o.rendererService,
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
	language        contexts.Language
	actions         pageactions.Actions
	filter          filters.Filter
	userGroups      []*group.Group
	canUpdate       bool
	canDelete       bool
	rendererService renderer.Renderer
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Overview.
func New(rendererService renderer.Renderer) *Overview {
	return &Overview{
		rendererService: rendererService,
	}
}
