package runner

import (
	"github.com/espal-digital-development/espal-core/validators/forms/account/login"
	"github.com/espal-digital-development/espal-core/validators/forms/account/password/change"
	"github.com/espal-digital-development/espal-core/validators/forms/account/password/forgot"
	"github.com/espal-digital-development/espal-core/validators/forms/account/password/recovery"
	"github.com/espal-digital-development/espal-core/validators/forms/account/register"
	"github.com/espal-digital-development/espal-core/validators/forms/auth"
)

type forms struct {
	admin *formsAdmin

	auth             auth.Factory
	login            login.Factory
	passwordChange   change.Factory
	passwordForgot   forgot.Factory
	passwordRecovery recovery.Factory
	register         register.Factory
}

func (r *Runner) forms() {
	r.formValidators.auth = auth.New(r.services.validators, r.stores.user)
	r.formValidators.login = login.New(r.services.validators, r.stores.user)
	r.formValidators.passwordChange = change.New(r.services.validators)
	r.formValidators.passwordForgot = forgot.New(r.services.validators)
	r.formValidators.passwordRecovery = recovery.New(r.services.validators, r.stores.user)
	r.formValidators.register = register.New(r.services.validators, r.stores.user)

	r.formsAdmin()
}
