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

func (r *Runner) routesAdminSite() error {
	createUpdatePage := createupdatepage.New(r.services.renderer)
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Site", overview.New(r.stores.site, overviewpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Site/View", view.New(r.repositories.languages, r.repositories.countries, r.repositories.currencies, r.stores.site, viewpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Site/Create", createupdate.New(r.services.entityMutators, r.stores.site, r.formValidators.admin.siteCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Site/Update", createupdate.New(r.services.entityMutators, r.stores.site, r.formValidators.admin.siteCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Site/ToggleOnline", toggleonline.New(r.repositories.regularExpressions, r.stores.site)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Site/Search", search.New(r.stores.site)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Site/Delete", delete.New(r.repositories.regularExpressions, r.stores.site)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
