package runner

import (
	"github.com/espal-digital-development/espal-core/storage/filesystem"
	"github.com/espal-digital-development/espal-core/storage/memory"
	"github.com/juju/errors"
)

func (r *Runner) assets() error {
	// TODO :: 77777 :: These 2 storages should be compatible with external service in the future (like S3 and so)
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
	err = r.storages.assetsPublicRootFiles.LoadAllFromPath(r.services.config.PublicRootFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsImages = memory.New()
	err = r.storages.assetsImages.LoadAllFromPath(r.services.config.ImagesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsStylesheets = memory.New()
	err = r.storages.assetsStylesheets.LoadAllFromPath(r.services.config.StylesheetsAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsJavaScript = memory.New()
	err = r.storages.assetsJavaScript.LoadAllFromPath(r.services.config.JavaScriptAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	if err := r.assetsFromModules(); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (r *Runner) assetsFromModules() error {
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
