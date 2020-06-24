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
func (httpContext *HTTPContext) GetDomain() Domain {
	return httpContext.domain
}

// GetSite returns the Site for the current route.
func (httpContext *HTTPContext) GetSite() Site {
	return httpContext.site
}

// GetLanguage returns the relevant Language for this request.
func (httpContext *HTTPContext) GetLanguage() (Language, error) {
	if httpContext.language != nil {
		return httpContext.language, nil
	}
	var err error
	user, ok, err := httpContext.GetUser()
	if err != nil {
		return nil, errors.Trace(err)
	}
	if ok {
		httpContext.language, err = httpContext.languagesRepository.ByID(user.Language())
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	if httpContext.language == nil && httpContext.GetDomain().Language() != nil {
		httpContext.language, err = httpContext.languagesRepository.ByID(*httpContext.GetDomain().Language())
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	if httpContext.language == nil && httpContext.GetSite().Language() != nil {
		httpContext.language, err = httpContext.languagesRepository.ByID(*httpContext.GetSite().Language())
		return httpContext.language, errors.Trace(err)
	}
	if httpContext.language != nil && !httpContext.configService.LanguageIsAvailable(httpContext.language.Code()) {
		httpContext.language, err = httpContext.languagesRepository.ByCode(httpContext.configService.DefaultLanguage())
		return httpContext.language, errors.Trace(err)
	}
	if httpContext.language != nil {
		return httpContext.language, nil
	}
	return nil, errors.Errorf("language not found")
}

// Translate is a shortcut to the passed translations service.
func (httpContext *HTTPContext) Translate(key string) string {
	language, err := httpContext.GetLanguage()
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	return httpContext.translationsRepository.Singular(language.ID(), key)
}

// TranslatePlural is a shortcut to the passed translations service.
func (httpContext *HTTPContext) TranslatePlural(key string) string {
	language, err := httpContext.GetLanguage()
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	return httpContext.translationsRepository.Plural(language.ID(), key)
}

// GetSlugMappedURL returns the slug when the route is a slug.
func (httpContext *HTTPContext) GetSlugMappedURL() string {
	return httpContext.slugMappedURL
}

// TODO :: 77 Maybe it's better to give it the Slug entity itself?

// SetSlugMappedURL sets the slug URL when the route came in as a slug.
// What this means is when a slug route was called (e.g. "/Inloggen")
// and it internally forwards to "/Login" the current route this function
// is called from will return "/Inloggen" to know the slug.
func (httpContext *HTTPContext) SetSlugMappedURL(slugMappedURL string) {
	httpContext.slugMappedURL = slugMappedURL
}
