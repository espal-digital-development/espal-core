package runner

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/routing/routes/api/v1/account/login"
	"github.com/espal-digital-development/espal-core/routing/routes/api/v1/account/overview"
	"github.com/juju/errors"
)

func (r *Runner) routesAPI() error {
	if err := r.services.router.RegisterRoute("/API/V1/Account", overview.New(r.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/API/V1/Login", login.New(r.stores.user)); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/API/V1/Account/Register", &apiEndPointNotImplemented{}); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/API/V1/ForgotPassword", &apiEndPointNotImplemented{}); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/API/V1/ForgotPasswordSucceeded", &apiEndPointNotImplemented{}); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/API/V1/PasswordRecovery", &apiEndPointNotImplemented{}); err != nil {
		return errors.Trace(err)
	}
	return nil
}

type apiEndPointNotImplemented struct {
}

func (a *apiEndPointNotImplemented) Handle(context contexts.Context) {
	if _, err := context.WriteString("This endpoint is not implemented yet."); err != nil {
		context.SetStatusCode(http.StatusInternalServerError)
	}
}
