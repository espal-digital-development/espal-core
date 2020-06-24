package runner

import (
	dashboardpage "github.com/espal-digital-development/espal-core/pages/admin/dashboard"
	"github.com/espal-digital-development/espal-core/routing/routes/admin/dashboard"
	"github.com/juju/errors"
)

func (runner *Runner) routesAdminCore() error {
	if err := runner.services.router.RegisterRoute(runner.services.config.AdminURL(), dashboard.New(dashboardpage.New())); err != nil {
		return errors.Trace(err)
	}
	return nil
}
