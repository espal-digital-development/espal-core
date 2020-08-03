package runner

import (
	"github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/juju/errors"
)

func (r *Runner) themes() error {
	defaultTheme := r.repositories.themes.NewTheme("base")
	// TODO :: 777777 Register and implement all other root-core views here
	authView := themes.NewView("auth")
	if err := defaultTheme.SetView(authView); err != nil {
		return errors.Trace(err)
	}

	if err := r.repositories.themes.Register(defaultTheme); err != nil {
		return errors.Trace(err)
	}

	for k := range r.modulesRegistry {
		moduleThemes, err := r.modulesRegistry[k].GetThemes()
		if err != nil {
			return errors.Trace(err)
		}
		if moduleThemes == nil {
			continue
		}
		for _, theme := range moduleThemes.Themes() {
			if err := r.repositories.themes.Register(theme); err != nil {
				return errors.Trace(err)
			}
		}
		for themeCode, views := range moduleThemes.Views() {
			theme, err := r.repositories.themes.GetTheme(themeCode)
			if err != nil {
				return errors.Trace(err)
			}
			for k2 := range views {
				if err := theme.SetView(views[k2]); err != nil {
					return errors.Trace(err)
				}
			}
		}
	}

	return nil
}
