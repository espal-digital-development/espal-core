package runner

import (
	"github.com/espal-digital-development/espal-core/adminmenu"
	servererrorpage "github.com/espal-digital-development/espal-core/pages/servererror"
	"github.com/espal-digital-development/espal-core/routing/router"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/template/renderer"
)

func (runner *Runner) router() {
	runner.services.renderer = renderer.New(runner.repositories.languages, runner.repositories.countries, runner.repositories.translations, runner.services.logger)
	runner.services.adminMenu = adminmenu.New(runner.services.config, runner.databases.selecter, runner.repositories.translations, runner.repositories.userRights)
	runner.services.contexts = contexts.New(runner.services.config, runner.services.logger, runner.repositories.languages, runner.repositories.translations, runner.services.sessions, runner.services.adminMenu, runner.services.renderer, runner.stores.user, servererrorpage.New())
	runner.services.router = router.New(runner.services.config, runner.services.logger, runner.services.contexts, runner.stores.domain, runner.stores.site, runner.stores.slug)
}
