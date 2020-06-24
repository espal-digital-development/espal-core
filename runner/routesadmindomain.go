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

func (runner *Runner) routesAdminDomain() error {
	createUpdatePage := createupdatepage.New(runner.services.renderer)
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Domain", overview.New(runner.stores.domain, overviewpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Domain/View", view.New(runner.repositories.languages, runner.stores.domain, viewpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Domain/Create", createupdate.New(runner.services.entityMutators, runner.stores.domain, runner.formValidators.admin.domainCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Domain/Update", createupdate.New(runner.services.entityMutators, runner.stores.domain, runner.formValidators.admin.domainCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Domain/ToggleActive", toggleactive.New(runner.repositories.regularExpressions, runner.stores.domain)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/Domain/Delete", delete.New(runner.repositories.regularExpressions, runner.stores.domain)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
