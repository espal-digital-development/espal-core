package runner

import (
	createupdatepage "github.com/espal-digital-development/espal-core/pages/admin/user/contact/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/contact/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/contact/delete"
	"github.com/juju/errors"
)

func (runner *Runner) routesAdminUserContact() error {
	createUpdatePage := createupdatepage.New(runner.services.renderer)
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Contact/Create", createupdate.New(runner.services.entityMutators, runner.stores.user, runner.stores.userContact, runner.formValidators.admin.userContactCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Contact/Update", createupdate.New(runner.services.entityMutators, runner.stores.user, runner.stores.userContact, runner.formValidators.admin.userContactCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Contact/Delete", delete.New(runner.repositories.regularExpressions, runner.stores.userContact)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
