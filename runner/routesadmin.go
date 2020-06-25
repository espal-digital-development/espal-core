package runner

import (
	"github.com/juju/errors"
)

func (r *Runner) routesAdmin() error {
	if err := r.routesAdminCore(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAdminUser(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAdminUserAddress(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAdminUserContact(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAdminUserGroup(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAdminSite(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAdminDomain(); err != nil {
		return errors.Trace(err)
	}
	return nil
}
