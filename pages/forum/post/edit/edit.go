package edit

import (
	"github.com/espal-digital-development/espal-core/pages/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
)

var _ Factory = &Edit{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, post *forum.Post) Template
}

// Edit page service.
type Edit struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (edit *Edit) NewPage(context contexts.Context, post *forum.Post) Template {
	page := &Page{
		post: post,
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
	post *forum.Post
}

// Render the page writing to the context.
func (page *Page) Render() {
	base.WritePageTemplate(page.GetCoreContext(), page)
}

// New returns a new instance of Edit.
func New() *Edit {
	return &Edit{}
}
