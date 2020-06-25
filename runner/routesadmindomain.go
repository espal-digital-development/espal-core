package runner

import (
	createupdatepage "github.com/espal-digital-development/espal-core/pages/admin/domain/createupdate"
	overviewpage "github.com/espal-digital-development/espal-core/pages/admin/domain/overview"
	viewpage "github.com/espal-digital-development/espal-core/pages/admin/domain/view"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/domain/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/domain/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/domain/overview"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/domain/toggleactive"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/domain/view"
	"github.com/juju/errors"
)

func (r *Runner) routesAdminDomain() error {
	createUpdatePage := createupdatepage.New(r.services.renderer)
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Domain", overview.New(r.stores.domain, overviewpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Domain/View", view.New(r.repositories.languages, r.stores.domain, viewpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Domain/Create", createupdate.New(r.services.entityMutators, r.stores.domain, r.formValidators.admin.domainCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Domain/Update", createupdate.New(r.services.entityMutators, r.stores.domain, r.formValidators.admin.domainCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Domain/ToggleActive", toggleactive.New(r.repositories.regularExpressions, r.stores.domain)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/Domain/Delete", delete.New(r.repositories.regularExpressions, r.stores.domain)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
