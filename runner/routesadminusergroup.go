package runner

import (
	overviewpage "github.com/espal-digital-development/espal-core/pages/admin/user/group/overview"
	viewpage "github.com/espal-digital-development/espal-core/pages/admin/user/group/view"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/group/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/group/overview"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/group/toggleactive"
	translationcreateupdate "github.com/espal-digital-development/espal-core/routing/routes/admin/user/group/translation/createupdate"
	translationdelete "github.com/espal-digital-development/espal-core/routing/routes/admin/user/group/translation/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/group/userrightupdate"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/user/group/view"
	"github.com/juju/errors"
)

func (runner *Runner) routesAdminUserGroup() error {
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup", overview.New(runner.stores.userGroup, overviewpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup/View", view.New(runner.repositories.userRights, runner.stores.userGroup, viewpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup/UserRights/Update", userrightupdate.New(runner.repositories.regularExpressions, runner.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup/Translations/Create", translationcreateupdate.New(runner.repositories.regularExpressions, runner.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup/Translations/Update", translationcreateupdate.New(runner.repositories.regularExpressions, runner.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup/Translations/Delete", translationdelete.New(runner.repositories.regularExpressions, runner.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup/ToggleActive", toggleactive.New(runner.repositories.regularExpressions, runner.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL()+"/UserGroup/Delete", delete.New(runner.repositories.regularExpressions, runner.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
