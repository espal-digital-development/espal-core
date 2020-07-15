package assethandler

import (
	"mime"
	"strings"

	"github.com/juju/errors"
	pngquant "github.com/yusukebe/go-pngquant"
)

func (h *AssetHandler) registerImagesRoutes() error {
	allowedExtensions := map[string]bool{
		"jpg":  true,
		"jpeg": true,
		"png":  true,
		"gif":  true,
		"webp": true,
		"svg":  true,
	}
	var loopErr error
	err := h.imagesStorage.Iterate(func(path string, data []byte, err error) bool {
		if err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		fileNameParts := strings.Split(path, ".")
		if len(fileNameParts) <= 1 {
			return true
		}
		extension := fileNameParts[len(fileNameParts)-1]
		if _, ok := allowedExtensions[extension]; !ok {
			return true
		}

		mimeType := mime.TypeByExtension("." + extension)

		switch mimeType {
		case "image/png":
			shrunkenSizeInBytes, err := pngquant.CompressBytes(data, "1")
			if err != nil {
				loopErr = errors.Trace(err)
				return false
			}
			if len(shrunkenSizeInBytes) < len(data) {
				data = shrunkenSizeInBytes
			}
			// case "image/jpeg":
			// 	// TODO :: 777777 jpegoptim (wrapper or cmd)
			// case "image/gif":
			// 	// TODO :: 777777 gifsicle (wrapper or cmd)
			// case "image/svg+xml":
			// 	// TODO :: 777777 svgo (wrapper or cmd)
		}

		var brotliData []byte
		if h.configService.AssetsBrotli() {
			brotliData, err = h.convertToBrotli(data)
			if err != nil {
				loopErr = errors.Trace(err)
				return false
			}
		}

		var gzipData []byte
		if h.configService.AssetsGZip() {
			gzipData, err = h.convertToGzip(data)
			if err != nil {
				loopErr = errors.Trace(err)
				return false
			}
		}

		err = h.routerService.RegisterRoute("/i/"+path, &route{
			data:        data,
			brotliData:  brotliData,
			gzipData:    gzipData,
			contentType: mimeType,
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
