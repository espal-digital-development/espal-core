package runner

import (
	domaincreateupdate "github.com/espal-digital-development/espal-core/validators/forms/admin/domain/createupdate"
	sitecreateupdate "github.com/espal-digital-development/espal-core/validators/forms/admin/site/createupdate"
	useraddresscreateupdate "github.com/espal-digital-development/espal-core/validators/forms/admin/user/address/createupdate"
	usercontactcreateupdate "github.com/espal-digital-development/espal-core/validators/forms/admin/user/contact/createupdate"
	usercreateupdate "github.com/espal-digital-development/espal-core/validators/forms/admin/user/createupdate"
)

type formsAdmin struct {
	domainCreateUpdate      domaincreateupdate.Factory
	siteCreateUpdate        sitecreateupdate.Factory
	userCreateUpdate        usercreateupdate.Factory
	userAddressCreateUpdate useraddresscreateupdate.Factory
	userContactCreateUpdate usercontactcreateupdate.Factory
}

func (runner *Runner) formsAdmin() {
	runner.formValidators.admin = &formsAdmin{}
	runner.formValidators.admin.domainCreateUpdate = domaincreateupdate.New(runner.services.validators, runner.repositories.translations, runner.services.config, runner.stores.site)
	runner.formValidators.admin.siteCreateUpdate = sitecreateupdate.New(runner.services.validators, runner.repositories.translations)
	runner.formValidators.admin.userCreateUpdate = usercreateupdate.New(runner.services.validators, runner.repositories.translations, runner.storages.assetsPublicFiles, runner.stores.userAddress, runner.stores.userContact)
	runner.formValidators.admin.userAddressCreateUpdate = useraddresscreateupdate.New(runner.services.validators)
	runner.formValidators.admin.userContactCreateUpdate = usercontactcreateupdate.New(runner.services.validators, runner.services.config, runner.stores.userContact)
}
