package runner

import (
	createupdatepage "github.com/espal-digital-development/espal-core/pages/admin/user/createupdate"
	overviewpage "github.com/espal-digital-development/espal-core/pages/admin/user/overview"
	viewpage "github.com/espal-digital-development/espal-core/pages/admin/user/view"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/createupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/overview"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/removeavatar"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/search"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/toggleactive"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/view"
	"github.com/juju/errors"
)

func (r *Runner) routesAdminUser() error {
	createUpdatePage := createupdatepage.New(r.services.renderer)
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User", overview.New(r.stores.user, overviewpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/View", view.New(r.stores.user, r.stores.userAddress, r.stores.userContact, viewpage.New(r.services.renderer, r.stores.user, r.stores.userContact))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Create", createupdate.New(r.services.config, r.services.assetHandler, r.services.entityMutators, r.storages.assetsPublicFiles, r.stores.user, r.stores.userAddress, r.formValidators.admin.userCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Update", createupdate.New(r.services.config, r.services.assetHandler, r.services.entityMutators, r.storages.assetsPublicFiles, r.stores.user, r.stores.userAddress, r.formValidators.admin.userCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/ToggleActive", toggleactive.New(r.repositories.regularExpressions, r.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Search", search.New(r.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/Delete", delete.New(r.repositories.regularExpressions, r.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/User/RemoveAvatar", removeavatar.New(r.services.config, r.storages.assetsPublicFiles, r.stores.user)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
