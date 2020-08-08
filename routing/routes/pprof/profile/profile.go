package profile

import (
	"net/http"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

// Route processor.
type Route struct{}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	var err error
	var seconds int
	if secondsStrings := context.QueryValue("seconds"); secondsStrings != "" {
		seconds, err = strconv.Atoi(secondsStrings)
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	context.SetContentType("application/octet-stream")
	if err := pprof.StartCPUProfile(context); err != nil {
		context.SetContentType("text/plain; charset=utf-8")
		context.SetStatusCode(http.StatusInternalServerError)
		context.WriteString("Failed to enable CPU profiling:" + err.Error() + "\n")
		return
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	pprof.StopCPUProfile()
}

// New returns a new instance of Route.
func New() *Route {
	return &Route{}
}
