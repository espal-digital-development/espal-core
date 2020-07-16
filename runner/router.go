package runner

import (
	"github.com/espal-digital-development/espal-core/adminmenu"
	"github.com/espal-digital-development/espal-core/routing/router"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/template/renderer"
)

func (r *Runner) router() {
	r.services.renderer = renderer.New(r.repositories.languages, r.repositories.countries, r.repositories.translations,
		r.services.logger)
	r.services.adminMenu = adminmenu.New(r.services.config, r.databases.selecter, r.repositories.translations,
		r.repositories.userRights)
	r.services.contexts = contexts.New(r.services.config, r.services.logger, r.repositories.languages,
		r.repositories.translations, r.services.sessions, r.services.adminMenu, r.services.renderer, r.stores.user)
	r.services.router = router.New(r.services.config, r.services.logger, r.services.contexts, r.stores.domain,
		r.stores.site, r.stores.slug)
}
