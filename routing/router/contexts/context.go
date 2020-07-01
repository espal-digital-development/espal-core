package contexts

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/adminmenu"
	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/sessions"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/juju/errors"
)

var _ Context = &HTTPContext{}

// Entity type interface.
type Entity interface {
	ID() string
	IsUpdated() bool
	CreatedByID() string
	UpdatedByID() *string
	CreatedByFirstName() *string
	CreatedBySurname() *string
	UpdatedByFirstName() *string
	UpdatedBySurname() *string
}

// Language type interface.
type Language interface {
	ID() uint16
	Code() string
	Translate(uint16) string
}

// Domain entity interface.
type Domain interface {
	ID() string
	Host() string
	Active() bool
	SiteID() string
	Language() *uint16
	HostWithProtocol() string
	HostWithProtocolAndWWW() string
}

// Site entity interface.
type Site interface {
	ID() string
	Online() bool
	Language() *uint16
}

// Context represents an object that holds all unique request information.
type Context interface {
	RenderStatusContext
	AuthenticationContext
	FlashMessageContext
	RequestContext
	FormContext
	AdminContext
	RenderContext

	GetDomain() Domain
	GetSite() Site
	GetLanguage() (Language, error)
	Translate(string) string
	TranslatePlural(string) string
	GetSlugMappedURL() string
	SetSlugMappedURL(slugMappedURL string)
	GetRequestMethod() string
}

// HTTPContext holds all unique request information.
type HTTPContext struct {
	configService          config.Config
	loggerService          logger.Loggable
	languagesRepository    languages.Repository
	translationsRepository translations.Repository
	sessionsFactory        sessions.Factory
	adminMenuService       adminmenu.Menu
	rendererService        renderer.Renderer
	userStore              user.Store
	serverError            ServerError

	request        *http.Request
	responseWriter http.ResponseWriter

	domain Domain
	site   Site

	httpStatusCode int
	formIsParsed   bool

	session            sessions.Session
	hasFlashMessage    bool
	flashMessageBuffer sessions.Message
	user               *user.User
	language           Language
	slugMappedURL      string
}

// GetDomain returns the Domain for the current route.
func (c *HTTPContext) GetDomain() Domain {
	return c.domain
}

// GetSite returns the Site for the current route.
func (c *HTTPContext) GetSite() Site {
	return c.site
}

// GetLanguage returns the relevant Language for this request.
func (c *HTTPContext) GetLanguage() (Language, error) {
	if c.language != nil {
		return c.language, nil
	}
	var err error
	user, ok, err := c.GetUser()
	if err != nil {
		return nil, errors.Trace(err)
	}
	if ok {
		c.language, err = c.languagesRepository.ByID(user.Language())
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	if c.language == nil && c.GetDomain().Language() != nil {
		c.language, err = c.languagesRepository.ByID(*c.GetDomain().Language())
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	if c.language == nil && c.GetSite().Language() != nil {
		c.language, err = c.languagesRepository.ByID(*c.GetSite().Language())
		return c.language, errors.Trace(err)
	}
	if c.language != nil && !c.configService.LanguageIsAvailable(c.language.Code()) {
		c.language, err = c.languagesRepository.ByCode(c.configService.DefaultLanguage())
		return c.language, errors.Trace(err)
	}
	if c.language != nil {
		return c.language, nil
	}
	return nil, errors.Errorf("language not found")
}

// Translate is a shortcut to the passed translations service.
func (c *HTTPContext) Translate(key string) string {
	language, err := c.GetLanguage()
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	return c.translationsRepository.Singular(language.ID(), key)
}

// TranslatePlural is a shortcut to the passed translations service.
func (c *HTTPContext) TranslatePlural(key string) string {
	language, err := c.GetLanguage()
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	return c.translationsRepository.Plural(language.ID(), key)
}

// GetSlugMappedURL returns the slug when the route is a slug.
func (c *HTTPContext) GetSlugMappedURL() string {
	return c.slugMappedURL
}

// TODO :: 77 Maybe it's better to give it the Slug entity itself?

// SetSlugMappedURL sets the slug URL when the route came in as a slug.
// What this means is when a slug route was called (e.g. "/Inloggen")
// and it internally forwards to "/Login" the current route this function
// is called from will return "/Inloggen" to know the slug.
func (c *HTTPContext) SetSlugMappedURL(slugMappedURL string) {
	c.slugMappedURL = slugMappedURL
}

// GetRequestMethod get's the request method.
func (c *HTTPContext) GetRequestMethod() string {
	return c.request.Method
}
