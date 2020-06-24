package runner

import (
	createupdatepage "github.com/espal-digital-development/espal-core/pages/admin/site/createupdate"
	overviewpage "github.com/espal-digital-development/espal-core/pages/admin/site/overview"
	viewpage "github.com/espal-digital-development/espal-core/pages/admin/site/view"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/site/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/site/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/site/overview"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/site/search"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/site/toggleonline"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/site/view"
	"github.com/juju/errors"
)

func (runner *Runner) routesAdminSite() error {
	createUpdatePage := createupdatepage.New(runner.services.renderer)
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Site", overview.New(runner.stores.site, overviewpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Site/View", view.New(runner.repositories.languages, runner.repositories.countries, runner.repositories.currencies, runner.stores.site, viewpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Site/Create", createupdate.New(runner.services.entityMutators, runner.stores.site, runner.formValidators.admin.siteCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Site/Update", createupdate.New(runner.services.entityMutators, runner.stores.site, runner.formValidators.admin.siteCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Site/ToggleOnline", toggleonline.New(runner.repositories.regularExpressions, runner.stores.site)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Site/Search", search.New(runner.stores.site)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Site/Delete", delete.New(runner.repositories.regularExpressions, runner.stores.site)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
