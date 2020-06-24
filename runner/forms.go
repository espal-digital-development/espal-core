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

func (runner *Runner) forms() {
	runner.formValidators.auth = auth.New(runner.services.validators, runner.stores.user)
	runner.formValidators.login = login.New(runner.services.validators, runner.stores.user)
	runner.formValidators.passwordChange = change.New(runner.services.validators)
	runner.formValidators.passwordForgot = forgot.New(runner.services.validators)
	runner.formValidators.passwordRecovery = recovery.New(runner.services.validators, runner.stores.user)
	runner.formValidators.register = register.New(runner.services.validators, runner.stores.user)

	runner.formsAdmin()
}
