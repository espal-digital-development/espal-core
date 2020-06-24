package runner

import (
	catalogpage "github.com/espal-digital-development/espal-core/pages/catalog"
	"github.com/espal-digital-development/espal-core/routing/routes/catalog"
	"github.com/juju/errors"
)

func (runner *Runner) routesCatalog() error {
	if err := runner.services.router.RegisterRoute("/Catalog", catalog.New(catalogpage.New())); err != nil {
		return errors.Trace(err)
	}
	return nil
}
