package runner

import (
	"github.com/juju/errors"
)

func (runner *Runner) routesAdmin() error {
	if err := runner.routesAdminCore(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAdminUser(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAdminUserAddress(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAdminUserContact(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAdminUserGroup(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAdminSite(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAdminDomain(); err != nil {
		return errors.Trace(err)
	}
	return nil
}
