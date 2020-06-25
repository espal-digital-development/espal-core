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
)

var _ Factory = &Contexts{}

// ServerError represents a spawner model.
type ServerError interface {
	RenderPage(context Context, title string, message string)
}

// Factory represents an object that facilitates the spawning of new Context objects.
type Factory interface {
	NewContext(request *http.Request, responseWriter http.ResponseWriter, domain Domain, site Site) Context
}

// Contexts facilitates the spawning of new Context objects.
type Contexts struct {
	configService          config.Config
	loggerService          logger.Loggable
	languagesRepository    languages.Repository
	translationsRepository translations.Repository
	sessionsFactory        sessions.Factory
	adminMenuService       adminmenu.Menu
	rendererService        renderer.Renderer
	userStore              user.Store
	serverError            ServerError
}

// NewContext returns a new instance of Context based on the given request- and router- information.
func (c *Contexts) NewContext(request *http.Request, responseWriter http.ResponseWriter, domain Domain, site Site) Context {
	context := &HTTPContext{
		configService:          c.configService,
		loggerService:          c.loggerService,
		languagesRepository:    c.languagesRepository,
		translationsRepository: c.translationsRepository,
		sessionsFactory:        c.sessionsFactory,
		adminMenuService:       c.adminMenuService,
		rendererService:        c.rendererService,
		userStore:              c.userStore,
		serverError:            c.serverError,

		request:        request,
		responseWriter: responseWriter,
		domain:         domain,
		site:           site,
	}
	context.SetContentType("text/html")
	return context
}

// New returns a new instance of Contexts.
func New(configService config.Config, loggerService logger.Loggable, languagesRepository languages.Repository, translationsRepository translations.Repository, sessionsFactory sessions.Factory, adminMenuService adminmenu.Menu, rendererService renderer.Renderer, userStore user.Store, serverError ServerError) *Contexts {
	return &Contexts{
		configService:          configService,
		loggerService:          loggerService,
		languagesRepository:    languagesRepository,
		translationsRepository: translationsRepository,
		sessionsFactory:        sessionsFactory,
		adminMenuService:       adminMenuService,
		rendererService:        rendererService,
		userStore:              userStore,
		serverError:            serverError,
	}
}
