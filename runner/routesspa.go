package runner

import (
	spapage "github.com/espal-digital-development/espal-core/pages/spa"
	"github.com/espal-digital-development/espal-core/routing/routes/spa"
	"github.com/juju/errors"
)

func (r *Runner) routesSPA() error {
	if err := r.services.router.RegisterRoute("/Spa", spa.New(spapage.New())); err != nil {
		return errors.Trace(err)
	}
	return nil
}
