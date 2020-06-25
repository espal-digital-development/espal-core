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

func (r *Runner) formsAdmin() {
	r.formValidators.admin = &formsAdmin{}
	r.formValidators.admin.domainCreateUpdate = domaincreateupdate.New(r.services.validators, r.repositories.translations, r.services.config, r.stores.site)
	r.formValidators.admin.siteCreateUpdate = sitecreateupdate.New(r.services.validators, r.repositories.translations)
	r.formValidators.admin.userCreateUpdate = usercreateupdate.New(r.services.validators, r.repositories.translations, r.storages.assetsPublicFiles, r.stores.userAddress, r.stores.userContact)
	r.formValidators.admin.userAddressCreateUpdate = useraddresscreateupdate.New(r.services.validators)
	r.formValidators.admin.userContactCreateUpdate = usercontactcreateupdate.New(r.services.validators, r.services.config, r.stores.userContact)
}
