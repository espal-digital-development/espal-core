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
func (r *route) Handle(context contexts.Context) {
	context.SetContentType(r.contentType)
	if r.cacheMaxAge != "" {
		context.SetHeader("Cache-Control", "max-age=3600")
	}
	if r.allowGzip && context.AcceptsEncoding("gzip") {
		context.SetHeader("Content-Encoding", "gzip")
		if _, err := context.Write(r.gzipData); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
		}
		return
	}
	if _, err := context.Write(r.data); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
	}
}
