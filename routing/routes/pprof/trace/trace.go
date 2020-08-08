package trace

import (
	"net/http"
	"runtime/trace"
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
	if seconds == 0 {
		seconds = 30
	}
	context.SetContentType("application/octet-stream")
	err = trace.Start(context)
	if err != nil {
		context.SetContentType("text/plain; charset=utf-8")
		context.SetStatusCode(http.StatusInternalServerError)
		context.WriteString("Failed to enable tracing: " + err.Error() + "\n")
		return
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	trace.Stop()
}

// New returns a new instance of Route.
func New() *Route {
	return &Route{}
}
