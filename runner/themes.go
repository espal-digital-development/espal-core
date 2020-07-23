package runner

import (
	"github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/juju/errors"
)

func (r *Runner) themes() error {
	defaultTheme := r.repositories.themes.NewTheme("base")
	authView := themes.NewView()

	if err := defaultTheme.AddView(authView); err != nil {
		return errors.Trace(err)
	}

	if err := r.repositories.themes.Register(defaultTheme); err != nil {
		return errors.Trace(err)
	}

	// TODO :: 777777 How to associate the theme/views with the respective Route/Context?
	// 	And how smart bind variable data inside it?

	return nil
}
