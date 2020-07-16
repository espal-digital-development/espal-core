package overview

import (
	"runtime/pprof"
	"strconv"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

// Route processor.
type Route struct{}

// Handle route handler.
// nolint:errcheck
func (r *Route) Handle(context contexts.Context) {
	context.WriteString(`<html><head><title>Pprof</title></head><body><h1>Pprof</h1><ul>`)
	for _, profile := range pprof.Profiles() {
		context.WriteString(`<li><a href="`)
		context.WriteString(context.PprofURL())
		context.WriteString(`/`)
		context.WriteString(profile.Name())
		context.WriteString(`?debug=1">`)
		context.WriteString(profile.Name())
		context.WriteString(` (`)
		context.WriteString(strconv.Itoa(profile.Count()))
		context.WriteString(`)</a></li>`)
	}
	context.WriteString(`<li><a href="`)
	context.WriteString(context.PprofURL())
	context.WriteString(`/goroutine?debug=2">Full Goroutine Stack Dump</a></li>`)
	context.WriteString(`</ul></body></html>`)
}

// New returns a new instance of Route.
func New() *Route {
	return &Route{}
}
