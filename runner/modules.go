package runner

import (
	"strings"

	"github.com/espal-digital-development/espal-core/modules"
	"github.com/juju/errors"
)

// RegisterModule register the given module into the runner app.
func (r *Runner) RegisterModule(module modules.Modular) error {
	meta := module.GetMeta()
	if meta == nil {
		return errors.Errorf("module meta definition is not set")
	}
	if meta.UniqueIdentifier() == "" {
		return errors.Errorf("module is missing it's unique identifier")
	}
	if meta.Name() == "" {
		return errors.Errorf("module `%s` is missing it's name", meta.UniqueIdentifier())
	}
	if meta.Version() == "" {
		return errors.Errorf("module `%s` is missing it's version", meta.UniqueIdentifier())
	}
	if !r.reValidSemver.MatchString(meta.Version()) {
		return errors.Errorf("module `%s` version is not valid semVer", meta.UniqueIdentifier())
	}
	for k := range r.modulesRegistry {
		if r.modulesRegistry[k].GetMeta().UniqueIdentifier() == meta.UniqueIdentifier() {
			return errors.Errorf("module `%s` is already registered",
				meta.UniqueIdentifier())
		}
	}
	r.services.logger.Infof("Registered module `%s` v%s", meta.Name(), strings.TrimPrefix(meta.Version(), "v"))
	r.modulesRegistry = append(r.modulesRegistry, module)
	return nil
}
