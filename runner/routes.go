package runner

import (
	"github.com/espal-digital-development/espal-core/modules/routes"
	authpage "github.com/espal-digital-development/espal-core/pages/auth"
	"github.com/espal-digital-development/espal-core/routing/assethandler"
	"github.com/espal-digital-development/espal-core/routing/routes/auth"
	"github.com/espal-digital-development/espal-core/routing/routes/health"
	authform "github.com/espal-digital-development/espal-core/validators/forms/auth"
	"github.com/juju/errors"
	"github.com/tdewolff/minify"
)

func (r *Runner) routes() error {
	r.services.assetHandler = assethandler.New(r.services.config, r.services.router, minify.New(),
		r.storages.assetsStylesheets, r.storages.assetsJavaScript, r.storages.assetsImages,
		r.storages.assetsPublicFiles, r.storages.assetsPublicRootFiles)
	if err := r.services.assetHandler.RegisterAll(); err != nil {
		return errors.Trace(err)
	}

	if err := r.services.router.RegisterRoute("/health", health.New()); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/Auth", auth.New(authform.New(r.services.validators, r.stores.user),
		authpage.New())); err != nil {
		return errors.Trace(err)
	}
	if r.services.config.Pprof() {
		if err := r.routesPprof(); err != nil {
			return errors.Trace(err)
		}
	}

	for k := range r.modulesRegistry {
		moduleRoutes, err := r.modulesRegistry[k].GetRoutes()
		if err != nil {
			return errors.Trace(err)
		}
		if moduleRoutes == nil {
			continue
		}
		err = moduleRoutes.Iterate(func(path string, h routes.Handler) error {
			return errors.Trace(r.services.router.RegisterRoute(path, h))
		})
		if err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
