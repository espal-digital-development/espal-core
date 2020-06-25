package runner

import (
	createupdatepage "github.com/espal-digital-development/espal-core/pages/admin/user/contact/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/contact/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/contact/delete"
	"github.com/juju/errors"
)

func (r *Runner) routesAdminUserContact() error {
	createUpdatePage := createupdatepage.New(r.services.renderer)
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Contact/Create", createupdate.New(r.services.entityMutators, r.stores.user, r.stores.userContact, r.formValidators.admin.userContactCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Contact/Update", createupdate.New(r.services.entityMutators, r.stores.user, r.stores.userContact, r.formValidators.admin.userContactCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Contact/Delete", delete.New(r.repositories.regularExpressions, r.stores.userContact)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
