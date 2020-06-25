package assethandler

import (
	"strings"

	"github.com/juju/errors"
)

func (h *AssetHandler) registerJavaScriptsRoutes() error {
	var gzipData []byte
	var loopErr error
	h.javaScriptStorage.Iterate(func(path string, data []byte, err error) bool {
		if err != nil {
			loopErr = err
			return false
		}
		data, err = h.minifier.Bytes("application/javascript", data)
		if err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		if !h.configService.Development() {
			data = []byte(strings.Replace(string(data), "'use strict';", "", -1))
		}
		if h.configService.AssetsGZip() {
			gzipData, err = h.convertToGzip(data)
			if err != nil {
				loopErr = errors.Trace(err)
				return false
			}
		}
		err = h.routerService.RegisterRoute("/j/"+path, &route{
			data:        data,
			gzipData:    gzipData,
			contentType: "application/javascript",
			cacheMaxAge: h.configService.AssetsCacheMaxAge(),
			allowGzip:   h.configService.AssetsGZip(),
		})
		if err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		return true
	})
	return errors.Trace(loopErr)
}
