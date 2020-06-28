package runner

import (
	"github.com/espal-digital-development/espal-core/adminmenu"
	"github.com/espal-digital-development/espal-core/cachesynchronizer"
	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/database/entitymutators"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/espal-digital-development/espal-core/database/queryhelper"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/mailer"
	"github.com/espal-digital-development/espal-core/routing/assethandler"
	"github.com/espal-digital-development/espal-core/routing/router"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/sessions"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/espal-digital-development/espal-core/tokenpool"
	"github.com/espal-digital-development/espal-core/validators"
)

type services struct {
	config              config.Config
	logger              logger.Loggable
	mailer              mailer.Engine
	tokenPool           tokenpool.Pool
	cacheSynchronizer   cachesynchronizer.Synchronizer
	sessions            sessions.Factory
	entityMutators      entitymutators.Factory
	databaseQueryHelper queryhelper.Helper
	databaseFilters     filters.Factory
	validators          validators.Factory
	contexts            contexts.Factory
	router              router.Router
	assetHandler        assethandler.Handler
	renderer            renderer.Renderer
	adminMenu           adminmenu.Menu
}
