package base

import (
	"io"

	"github.com/espal-digital-development/espal-core/sessions"
)

// Context core context.
type Context interface {
	io.Writer
	IsLoggedIn() bool
	HasAdminAccess() bool
	HasPprofEnabled() bool
	AdminURL() string
	PprofURL() string
	Translate(string) string
	TranslatePlural(string) string
	HasFlashMessage() bool
	GetFlashMessage() sessions.Message
}

// Form interactive handling.
type Form interface {
	Errors() string
	Open() string
	Field(string) string
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

// HasAdminAccess returns an indicator if the user has administrator access.
func (page *Page) HasAdminAccess() bool {
	return page.coreContext.HasAdminAccess()
}

// HasPprofEnabled returns an indicator if the user has pprof access.
func (page *Page) HasPprofEnabled() bool {
	return page.coreContext.HasPprofEnabled()
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

// PprofURL returns the url prefix for visiting pprof area paths.
func (page *Page) PprofURL() string {
	return page.coreContext.PprofURL()
}

// HasFlashMessage returns an indicator if a flash message was set.
func (page *Page) HasFlashMessage() bool {
	return page.coreContext.HasFlashMessage()
}

// GetFlashMessage returns the set flash message.
func (page *Page) GetFlashMessage() sessions.Message {
	return page.coreContext.GetFlashMessage()
}
