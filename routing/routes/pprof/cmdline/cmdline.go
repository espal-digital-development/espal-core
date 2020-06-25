package cmdline

import (
	"os"
	"runtime/pprof"
	"strings"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

// Route processor.
type Route struct{}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	context.SetContentType("text/plain; charset=utf-8")
	if _, err := context.WriteString(strings.Join(os.Args, "\x00")); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	pprof.Lookup("name")
}

// New returns a new instance of Route.
func New() *Route {
	return &Route{}
}
