package runner

import (
	"strings"
	"time"

	"github.com/espal-digital-development/espal-core/modules"
	"github.com/juju/errors"
)

// RegisterModule register the given module into the runner app.
func (r *Runner) RegisterModule(module modules.Modular) error {
	start := time.Now()
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
	if !r.services.semver.Valid(meta.Version()) {
		return errors.Errorf("module `%s` version is not valid semVer", meta.UniqueIdentifier())
	}
	for k := range r.modulesRegistry {
		if r.modulesRegistry[k].GetMeta().UniqueIdentifier() == meta.UniqueIdentifier() {
			return errors.Errorf("module `%s` is already registered",
				meta.UniqueIdentifier())
		}
	}

	var err error
	var compatible bool

	switch {
	case meta.MinimumCompatibleCoreVersion() != "" && meta.MaximumCompatibleCoreVersion() != "":
		compatible, err = r.services.semver.InRange(r.version, meta.MinimumCompatibleCoreVersion(),
			meta.MaximumCompatibleCoreVersion())
		if err != nil {
			return errors.Trace(err)
		}
	case meta.MinimumCompatibleCoreVersion() != "":
		compatible, err = r.services.semver.GreaterThanOrEqual(r.version, meta.MinimumCompatibleCoreVersion())
		if err != nil {
			return errors.Trace(err)
		}
	case meta.MaximumCompatibleCoreVersion() != "":
		compatible, err = r.services.semver.SmallerThanOrEqual(r.version, meta.MinimumCompatibleCoreVersion())
		if err != nil {
			return errors.Trace(err)
		}
	}
	if !compatible {
		return errors.Errorf("module `%s` is not version-compatible with core version `%s`. Has to be between `%s` and `%s`",
			meta.UniqueIdentifier(), r.version, meta.MinimumCompatibleCoreVersion(),
			meta.MaximumCompatibleCoreVersion())
	}

	r.services.logger.Infof("Registered module `%s` v%s in %v",
		meta.Name(), strings.TrimPrefix(meta.Version(), "v"), time.Since(start))
	r.modulesRegistry = append(r.modulesRegistry, module)
	return nil
}
