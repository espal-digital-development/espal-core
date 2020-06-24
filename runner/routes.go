package runner

import (
	"github.com/espal-digital-development/espal-core/routing/assethandler"
	"github.com/juju/errors"
	"github.com/tdewolff/minify"
)

func (runner *Runner) routes() error {
	runner.services.assetHandler = assethandler.New(runner.services.config, runner.services.router, minify.New(), runner.storages.assetsStylesheets, runner.storages.assetsJavaScript, runner.storages.assetsImages, runner.storages.assetsPublicFiles, runner.storages.assetsPublicRootFiles)
	if err := runner.services.assetHandler.RegisterAll(); err != nil {
		return errors.Trace(err)
	}
	if runner.services.config.Pprof() {
		if err := runner.routesPprof(); err != nil {
			return errors.Trace(err)
		}
	}
	if err := runner.routesAPI(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesFrontend(); err != nil {
		return errors.Trace(err)
	}
	if err := runner.routesAdmin(); err != nil {
		return errors.Trace(err)
	}
	return nil
}
