package runner

import (
	"github.com/juju/errors"
)

func (runner *Runner) routesFrontend() error {
	if err := runner.routesCore(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAccount(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesCatalog(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesForum(); err != nil {
		return errors.Trace(err)
	}
	return nil
}
