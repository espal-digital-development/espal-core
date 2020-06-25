package runner

import (
	createupdatepage "github.com/espal-digital-development/espal-core/pages/admin/user/address/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/address/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/address/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/address/toggleactive"
	"github.com/juju/errors"
)

func (r *Runner) routesAdminUserAddress() error {
	createUpdatePage := createupdatepage.New(r.services.renderer)
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Address/Create", createupdate.New(r.services.entityMutators, r.stores.user, r.stores.userAddress, r.formValidators.admin.userAddressCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Address/Update", createupdate.New(r.services.entityMutators, r.stores.user, r.stores.userAddress, r.formValidators.admin.userAddressCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Address/ToggleActive", toggleactive.New(r.repositories.regularExpressions, r.stores.userAddress)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Address/Delete", delete.New(r.repositories.regularExpressions, r.stores.userAddress)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
