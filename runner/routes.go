package runner

import (
	"github.com/espal-digital-development/espal-core/routing/assethandler"
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
	if r.services.config.Pprof() {
		if err := r.routesPprof(); err != nil {
			return errors.Trace(err)
		}
	}
	if err := r.routesAPI(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesFrontend(); err != nil {
		return errors.Trace(err)
	}
	if err := r.routesAdmin(); err != nil {
		return errors.Trace(err)
	}
	return nil
}
