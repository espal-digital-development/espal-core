package posts

import (
	"net/http"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	pprofPrefix string
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	name := strings.TrimPrefix(context.Path(), r.pprofPrefix)

	context.SetContentType("text/plain; charset=utf-8")
	var debug int
	if debugString := context.QueryValue("debug"); debugString != "" {
		var err error
		debug, err = strconv.Atoi(debugString)
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	profiler := pprof.Lookup(name)
	if profiler == nil {
		context.SetStatusCode(http.StatusNotFound)
		context.WriteString("Unknown profile: " + name + "\n")
		return
	}
	var gc int
	if gcString := context.QueryValue("gc"); gcString != "" {
		var err error
		gc, err = strconv.Atoi(gcString)
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	if gc > 0 && name == "heap" {
		runtime.GC()
	}
	if err := profiler.WriteTo(context, debug); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
}

// New returns a new instance of Route.
func New(pprofPrefix string) *Route {
	return &Route{
		pprofPrefix: pprofPrefix,
	}
}
