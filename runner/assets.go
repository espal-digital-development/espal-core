package runner

import (
	"github.com/espal-digital-development/espal-core/storage/filesystem"
	"github.com/espal-digital-development/espal-core/storage/memory"
	"github.com/juju/errors"
)

func (r *Runner) assets() error {
	// TODO :: 77777 These 2 storages should be compatible with external service in the future (like S3 and so)
	var err error
	r.storages.assetsPrivateFiles, err = filesystem.New(r.services.config.PrivateFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsPublicFiles, err = filesystem.New(r.services.config.PublicFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}

	assetsPublicRootFiles := memory.New()
	assetsImages := memory.New()
	assetsStylesheets := memory.New()
	assetsJavaScript := memory.New()

	for k := range r.modulesRegistry {
		assets := r.modulesRegistry[k].GetAssets()
		if assets == nil {
			continue
		}
		if err := assets.SetPublicRootFiles(assetsPublicRootFiles); err != nil {
			return errors.Trace(err)
		}
		if err := assets.SetImages(assetsImages); err != nil {
			return errors.Trace(err)
		}
		if err := assets.SetStylesheets(assetsStylesheets); err != nil {
			return errors.Trace(err)
		}
		if err := assets.SetJavaScript(assetsJavaScript); err != nil {
			return errors.Trace(err)
		}
	}

	if err := assetsPublicRootFiles.LoadAllFromPath(r.services.config.PublicRootFilesAssetsPath()); err != nil {
		return errors.Trace(err)
	}
	if err := assetsImages.LoadAllFromPath(r.services.config.ImagesAssetsPath()); err != nil {
		return errors.Trace(err)
	}
	if err := assetsStylesheets.LoadAllFromPath(r.services.config.StylesheetsAssetsPath()); err != nil {
		return errors.Trace(err)
	}
	if err := assetsJavaScript.LoadAllFromPath(r.services.config.JavaScriptAssetsPath()); err != nil {
		return errors.Trace(err)
	}

	r.storages.assetsPublicRootFiles = assetsPublicRootFiles
	r.storages.assetsImages = assetsImages
	r.storages.assetsStylesheets = assetsStylesheets
	r.storages.assetsJavaScript = assetsJavaScript

	return nil
}
