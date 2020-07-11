package runner

import (
	"github.com/juju/errors"
)

func (r *Runner) routesFrontend() error {
	if err := r.routesCore(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAccount(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesCatalog(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesForum(); err != nil {
		return errors.Trace(err)
	}
	return nil
}
