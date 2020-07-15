package assethandler

import (
	"github.com/juju/errors"
)

func (h *AssetHandler) registerStylesheetsRoutes() error {
	var brotliData []byte
	var gzipData []byte
	var loopErr error
	err := h.stylesheetsStorage.Iterate(func(path string, data []byte, err error) bool {
		if err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		data, err = h.minifier.Bytes("text/css", data)
		if err != nil {
			loopErr = errors.Trace(err)
			return false
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
		err = h.routerService.RegisterRoute("/c/"+path, &route{
			data:        data,
			brotliData:  brotliData,
			gzipData:    gzipData,
			contentType: "text/css",
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
