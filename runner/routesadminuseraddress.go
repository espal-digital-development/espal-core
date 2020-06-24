package runner

import (
	createupdatepage "github.com/espal-digital-development/espal-core/pages/admin/user/address/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/address/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/address/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/address/toggleactive"
	"github.com/juju/errors"
)

func (runner *Runner) routesAdminUserAddress() error {
	createUpdatePage := createupdatepage.New(runner.services.renderer)
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Address/Create", createupdate.New(runner.services.entityMutators, runner.stores.user, runner.stores.userAddress, runner.formValidators.admin.userAddressCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Address/Update", createupdate.New(runner.services.entityMutators, runner.stores.user, runner.stores.userAddress, runner.formValidators.admin.userAddressCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Address/ToggleActive", toggleactive.New(runner.repositories.regularExpressions, runner.stores.userAddress)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Address/Delete", delete.New(runner.repositories.regularExpressions, runner.stores.userAddress)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
