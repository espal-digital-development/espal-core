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

func (r *Runner) routesAdminUserGroup() error {
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup", overview.New(r.stores.userGroup, overviewpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup/View", view.New(r.repositories.userRights, r.stores.userGroup, viewpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup/UserRights/Update", userrightupdate.New(r.repositories.regularExpressions, r.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup/Translations/Create", translationcreateupdate.New(r.repositories.regularExpressions, r.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup/Translations/Update", translationcreateupdate.New(r.repositories.regularExpressions, r.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup/Translations/Delete", translationdelete.New(r.repositories.regularExpressions, r.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup/ToggleActive", toggleactive.New(r.repositories.regularExpressions, r.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL()+"/UserGroup/Delete", delete.New(r.repositories.regularExpressions, r.stores.userGroup)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
