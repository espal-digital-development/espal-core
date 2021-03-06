package runner

import (
	"github.com/espal-digital-development/espal-core/routing/routes/pprof/cmdline"
	"github.com/espal-digital-development/espal-core/routing/routes/pprof/overview"
	"github.com/espal-digital-development/espal-core/routing/routes/pprof/posts"
	"github.com/espal-digital-development/espal-core/routing/routes/pprof/profile"
	"github.com/espal-digital-development/espal-core/routing/routes/pprof/symbol"
	"github.com/espal-digital-development/espal-core/routing/routes/pprof/trace"
	"github.com/juju/errors"
)

// TODO :: When online; some way to protect these routes? Accidentally having these on might expose a system. Just
// relying on the user deploying with the  right settings is not enough. Need to make sure that the system can never be
// exploited.
func (r *Runner) routesPprof() error {
	pprofPrefix := r.services.config.PprofURL() + "/"

	if err := r.services.router.RegisterRoute(r.services.config.PprofURL(), overview.New()); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(pprofPrefix+"cmdline", cmdline.New()); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(pprofPrefix+"profile", profile.New()); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(pprofPrefix+"symbol", symbol.New()); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute(pprofPrefix+"trace", trace.New()); err != nil {
		return errors.Trace(err)
	}
	for _, postPath := range []string{"goroutine", "threadcreate", "block", "mutex", "heap", "allocs"} {
		if err := r.services.router.RegisterRoute(pprofPrefix+postPath, posts.New(pprofPrefix)); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}
