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

func (runner *Runner) routesAdminUser() error {
	createUpdatePage := createupdatepage.New(runner.services.renderer)
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User", overview.New(runner.stores.user, overviewpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/View", view.New(runner.stores.user, runner.stores.userAddress, runner.stores.userContact, viewpage.New(runner.services.renderer, runner.stores.user, runner.stores.userContact))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Create", createupdate.New(runner.services.config, runner.services.assetHandler, runner.services.entityMutators, runner.storages.assetsPublicFiles, runner.stores.user, runner.stores.userAddress, runner.formValidators.admin.userCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Update", createupdate.New(runner.services.config, runner.services.assetHandler, runner.services.entityMutators, runner.storages.assetsPublicFiles, runner.stores.user, runner.stores.userAddress, runner.formValidators.admin.userCreateUpdate, createUpdatePage)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/ToggleActive", toggleactive.New(runner.repositories.regularExpressions, runner.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Search", search.New(runner.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/Delete", delete.New(runner.repositories.regularExpressions, runner.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/User/RemoveAvatar", removeavatar.New(runner.services.config, runner.storages.assetsPublicFiles, runner.stores.user)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
