package runner

import (
	dashboardpage "github.com/espal-digital-development/espal-core/pages/admin/dashboard"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/dashboard"
	"github.com/juju/errors"
)

func (r *Runner) routesAdminCore() error {
	if err := r.services.router.RegisterRoute(r.services.config.AdminURL(), dashboard.New(dashboardpage.New())); err != nil {
		return errors.Trace(err)
	}
	return nil
}
