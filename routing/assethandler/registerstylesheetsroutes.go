package assethandler

import (
	"github.com/juju/errors"
)

func (assetHandler *AssetHandler) registerStylesheetsRoutes() error {
	var gzipData []byte
	var loopErr error
	assetHandler.stylesheetsStorage.Iterate(func(path string, data []byte, err error) bool {
		if err != nil {
			loopErr = err
			return false
		}
		data, err = assetHandler.minifier.Bytes("text/css", data)
		if err != nil {
			loopErr = err
			return false
		}
		if assetHandler.configService.AssetsGZip() {
			gzipData, err = assetHandler.convertToGzip(data)
			if err != nil {
				loopErr = err
				return false
			}
		}
		err = assetHandler.routerService.RegisterRoute("/c/"+path, &route{
			data:        data,
			gzipData:    gzipData,
			contentType: "text/css",
			cacheMaxAge: assetHandler.configService.AssetsCacheMaxAge(),
			allowGzip:   assetHandler.configService.AssetsGZip(),
		})
		if err != nil {
			loopErr = err
			return false
		}

		return true
	})
	return errors.Trace(loopErr)
}
