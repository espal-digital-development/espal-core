package assethandler

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

type route struct {
	data        []byte
	gzipData    []byte
	contentType string
	cacheMaxAge string
	allowGzip   bool
}

// Handle processes the asset calls and outputs the bytes of the requested file.
func (route *route) Handle(context contexts.Context) {
	context.SetContentType(route.contentType)
	if route.cacheMaxAge != "" {
		context.SetHeader("Cache-Control", "max-age=3600")
	}
	if route.allowGzip && context.AcceptsEncoding("gzip") {
		context.SetHeader("Content-Encoding", "gzip")
		if _, err := context.Write(route.gzipData); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
		}
		return
	}
	if _, err := context.Write(route.data); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
	}
}
