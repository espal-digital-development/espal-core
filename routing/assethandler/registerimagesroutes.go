package assethandler

import (
	"mime"
	"strings"

	"github.com/juju/errors"
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
			loopErr = err
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

		// if image.MimeType == "image/png" {
		// 	// TODO :: 777777 Enable again when it works on Windows
		// 	// shrunkenSizeInBytes, err := pngquant.Crush(files[k2], pngquant.SPEED_SLOWEST)
		// 	// if err != nil {
		// 	// 	return errors.Trace(err)
		// 	// }
		// 	// if len(shrunkenSizeInBytes) < len(image.Data) {
		// 	// 	image.Data = shrunkenSizeInBytes
		// 	// }
		// } else if image.MimeType == "image/jpeg" {
		// 	// TODO :: jpegoptim (wrapper or cmd)
		// } else if image.MimeType == "image/gif" {
		// 	// TODO :: gifsicle (wrapper or cmd)
		// } else if image.MimeType == "image/svg+xml" {
		// 	// TODO :: svgo (wrapper or cmd)
		// }

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
