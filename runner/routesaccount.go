package runner

import (
	loginpage "github.com/espal-digital-development/espal-core/pages/account/login"
	overviewpage "github.com/espal-digital-development/espal-core/pages/account/overview"
	passwordchangepage "github.com/espal-digital-development/espal-core/pages/account/password/change"
	passwordforgotpage "github.com/espal-digital-development/espal-core/pages/account/password/forgot"
	passwordforgotsucceededpage "github.com/espal-digital-development/espal-core/pages/account/password/forgot/succeeded"
	passwordrecoverypage "github.com/espal-digital-development/espal-core/pages/account/password/recovery"
	registerpage "github.com/espal-digital-development/espal-core/pages/account/register"
	registersucceededpage "github.com/espal-digital-development/espal-core/pages/account/register/succeeded"
	"github.com/espal-digital-development/espal-core/routing/routes/account/activate"
	"github.com/espal-digital-development/espal-core/routing/routes/account/login"
	"github.com/espal-digital-development/espal-core/routing/routes/account/logout"
	"github.com/espal-digital-development/espal-core/routing/routes/account/overview"
	"github.com/espal-digital-development/espal-core/routing/routes/account/password/change"
	"github.com/espal-digital-development/espal-core/routing/routes/account/password/forgot"
	passwordforgotsucceeded "github.com/espal-digital-development/espal-core/routing/routes/account/password/forgot/succeeded"
	"github.com/espal-digital-development/espal-core/routing/routes/account/password/recovery"
	"github.com/espal-digital-development/espal-core/routing/routes/account/register"
	registersucceeded "github.com/espal-digital-development/espal-core/routing/routes/account/register/succeeded"
	"github.com/juju/errors"
)

func (runner *Runner) routesAccount() error {
	if err := runner.services.router.RegisterRoute("/Login", login.New(runner.formValidators.login, loginpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/Logout", logout.New()); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/RegisterAccount", register.New(runner.services.config, runner.services.mailer, runner.stores.user, runner.formValidators.register, registerpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/RegisterAccountSucceeded", registersucceeded.New(registersucceededpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/ActivateAccount", activate.New(runner.repositories.regularExpressions, runner.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/ForgotPassword", forgot.New(runner.services.config, runner.services.mailer, runner.stores.user, runner.formValidators.passwordForgot, passwordforgotpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/ForgotPasswordSucceeded", passwordforgotsucceeded.New(passwordforgotsucceededpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/PasswordRecovery", recovery.New(runner.services.config, runner.repositories.regularExpressions, runner.stores.user, runner.formValidators.passwordRecovery, passwordrecoverypage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/Account", overview.New(overviewpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/Account/ChangePassword", change.New(runner.services.config, runner.stores.user, runner.formValidators.passwordChange, passwordchangepage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	return nil
}
