package runner

import (
	authpage "github.com/espal-digital-development/espal-core/pages/auth"
	rootpage "github.com/espal-digital-development/espal-core/pages/root"
	"github.com/espal-digital-development/espal-core/routing/routes/auth"
	"github.com/espal-digital-development/espal-core/routing/routes/root"
	"github.com/juju/errors"
)

func (runner *Runner) routesCore() error {
	if err := runner.services.router.RegisterRoute("/", root.New(rootpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/Auth", auth.New(runner.formValidators.auth, authpage.New())); err != nil {
		return errors.Trace(err)
	}
	return nil
}
