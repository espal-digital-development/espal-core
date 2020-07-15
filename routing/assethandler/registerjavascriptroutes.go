package assethandler

import (
	"strings"

	"github.com/juju/errors"
)

func (h *AssetHandler) registerJavaScriptsRoutes() error {
	var brotliData []byte
	var gzipData []byte
	var loopErr error
	err := h.javaScriptStorage.Iterate(func(path string, data []byte, err error) bool {
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
		if h.configService.AssetsBrotli() {
			brotliData, err = h.convertToBrotli(data)
			if err != nil {
				loopErr = errors.Trace(err)
				return false
			}
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
			brotliData:  brotliData,
			gzipData:    gzipData,
			contentType: "application/javascript",
			cacheMaxAge: h.configService.AssetsCacheMaxAge(),
			allowBrotli: h.configService.AssetsBrotli(),
			allowGzip:   h.configService.AssetsGZip(),
		})
		if err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		return true
	})
	if err != nil {
		return errors.Trace(err)
	}
	return errors.Trace(loopErr)
}
