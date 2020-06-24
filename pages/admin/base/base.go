package base

import (
	"io"

	"github.com/espal-digital-development/espal-core/adminmenu"
	"github.com/espal-digital-development/espal-core/sessions"
)

// Context core context.
type Context interface {
	io.Writer
	IsLoggedIn() bool
	AdminURL() string
	Translate(string) string
	TranslatePlural(string) string
	HasFlashMessage() bool
	GetFlashMessage() sessions.Message
	HasUserRight(string) bool
	AdminMainMenu() []*adminmenu.Block
}

// Form interactive handling.
type Form interface {
	Errors() string
	Open() string
	Field(string) string
	CreateUpdateActions(string, string) string
	ContainsSelectSearch() bool
}

// Page template object.
type Page struct {
	coreContext Context
}

// SetCoreContext sets the basic context requirements of the page.
func (page *Page) SetCoreContext(context Context) {
	page.coreContext = context
}

// GetCoreContext returns the internal core context.
func (page *Page) GetCoreContext() Context {
	return page.coreContext
}

// IsLoggedIn returns an indicator if the user is logged in or not.
func (page *Page) IsLoggedIn() bool {
	return page.coreContext.IsLoggedIn()
}

// Translate translates the given key based on the language
// active in the current context.
func (page *Page) Translate(key string) string {
	return page.coreContext.Translate(key)
}

// TranslatePlural translates the given key based on the language
// active in the current context in plural.
func (page *Page) TranslatePlural(key string) string {
	return page.coreContext.TranslatePlural(key)
}

// AdminURL returns the url prefix for visiting admin area paths.
func (page *Page) AdminURL() string {
	return page.coreContext.AdminURL()
}

// HasFlashMessage returns an indicator if a flash message was set.
func (page *Page) HasFlashMessage() bool {
	return page.coreContext.HasFlashMessage()
}

// GetFlashMessage returns the set flash message.
func (page *Page) GetFlashMessage() sessions.Message {
	return page.coreContext.GetFlashMessage()
}

// HasUserRight returns if the current user (if logged in) has the userright.
func (page *Page) HasUserRight(userRight string) bool {
	return page.coreContext.HasUserRight(userRight)
}

// AdminMainMenu returns the rendered admin menu for the current user.
func (page *Page) AdminMainMenu() []*adminmenu.Block {
	return page.coreContext.AdminMainMenu()
}
