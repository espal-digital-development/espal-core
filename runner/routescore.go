package runner

import (
	authpage "github.com/espal-digital-development/espal-core/pages/auth"
	rootpage "github.com/espal-digital-development/espal-core/pages/root"
	"github.com/espal-digital-development/espal-core/routing/routes/auth"
	"github.com/espal-digital-development/espal-core/routing/routes/root"
	"github.com/juju/errors"
)

func (r *Runner) routesCore() error {
	if err := r.services.router.RegisterRoute("/", root.New(rootpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/Auth", auth.New(r.formValidators.auth, authpage.New())); err != nil {
		return errors.Trace(err)
	}
	return nil
}
