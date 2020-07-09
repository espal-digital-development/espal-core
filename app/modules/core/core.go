package core

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/espal-digital-development/espal-core/modules"
	"github.com/espal-digital-development/espal-core/modules/assets"
	"github.com/espal-digital-development/espal-core/modules/meta"
	"github.com/espal-digital-development/espal-core/modules/translations"
	"github.com/juju/errors"
)

// TODO :: 777777
// - Some kind of of registry system for smart binding
// - How to hook into existing functionality like Slugs and other Database/Repository functionality
//   - How to get functionality báck into the modules? Some kind of reverse registration injection with interface{}'s?
// - CompatibilityDefintion should describe what versions of the core app works with and
//   whether it colides with other functionality being present in the system through other modules

// New returns a new instance of Module.
func New() (*modules.Module, error) {
	_, filename, _, _ := runtime.Caller(0) // nolint:dogsled
	cwd := path.Dir(filename)

	meta, err := meta.New(&meta.Config{
		UniqueIdentifier:             "com.espal.core",
		Version:                      "0.0.1",
		MinimumCompatibleCoreVersion: "0.0.1",
		MaximumCompatibleCoreVersion: "",
		Name:                         "Espal Core",
		Author:                       "Espal Digital Development",
		Contact:                      "https://github.com/espal-digital-development",
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	// TODO :: 777777 :: Try assets and translations first (move global translations from the frontend in here)
	assets, err := assets.New(&assets.Config{
		PublicRootFilesPath: filepath.FromSlash(cwd + "/app/assets/files/root"),
		ImagesPath:          filepath.FromSlash(cwd + "/app/assets/images"),
		StylesheetsPath:     filepath.FromSlash(cwd + "/app/assets/css"),
		JavaScriptPath:      filepath.FromSlash(cwd + "/app/assets/js"),
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	translations, err := translations.New()
	if err != nil {
		return nil, errors.Trace(err)
	}
	m, err := modules.New(&modules.Config{
		MetaDefinition:       meta,
		AssetsProvider:       assets,
		TranslationsProvider: translations,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}
	return m, nil
}
