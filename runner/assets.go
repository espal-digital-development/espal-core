package runner

import (
	"github.com/espal-digital-development/espal-core/storage/filesystem"
	"github.com/espal-digital-development/espal-core/storage/memory"
	"github.com/juju/errors"
)

func (r *Runner) assets() error {
	var err error
	r.storages.assetsPrivateFiles, err = filesystem.New(r.services.config.PrivateFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsPublicFiles, err = filesystem.New(r.services.config.PublicFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}

	r.storages.assetsPublicRootFiles = memory.New()
	r.storages.assetsImages = memory.New()
	r.storages.assetsStylesheets = memory.New()
	r.storages.assetsJavaScript = memory.New()
	if err != nil {
		return errors.Trace(err)
	}

	for k := range r.modulesRegistry {
		assets := r.modulesRegistry[k].GetAssets()
		if assets == nil {
			continue
		}
		if err := assets.SetPublicRootFiles(r.storages.assetsPublicRootFiles); err != nil {
			return errors.Trace(err)
		}
		if err := assets.SetImages(r.storages.assetsImages); err != nil {
			return errors.Trace(err)
		}
		if err := assets.SetStylesheets(r.storages.assetsStylesheets); err != nil {
			return errors.Trace(err)
		}
		if err := assets.SetJavaScript(r.storages.assetsJavaScript); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
