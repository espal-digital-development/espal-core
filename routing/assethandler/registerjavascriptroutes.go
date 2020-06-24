package assethandler

import (
	"strings"

	"github.com/juju/errors"
)

func (assetHandler *AssetHandler) registerJavaScriptsRoutes() error {
	var gzipData []byte
	var loopErr error
	assetHandler.javaScriptStorage.Iterate(func(path string, data []byte, err error) bool {
		if err != nil {
			loopErr = err
			return false
		}
		data, err = assetHandler.minifier.Bytes("application/javascript", data)
		if err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		if !assetHandler.configService.Development() {
			data = []byte(strings.Replace(string(data), "'use strict';", "", -1))
		}
		if assetHandler.configService.AssetsGZip() {
			gzipData, err = assetHandler.convertToGzip(data)
			if err != nil {
				loopErr = errors.Trace(err)
				return false
			}
		}
		err = assetHandler.routerService.RegisterRoute("/j/"+path, &route{
			data:        data,
			gzipData:    gzipData,
			contentType: "application/javascript",
			cacheMaxAge: assetHandler.configService.AssetsCacheMaxAge(),
			allowGzip:   assetHandler.configService.AssetsGZip(),
		})
		if err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		return true
	})
	return errors.Trace(loopErr)
}
