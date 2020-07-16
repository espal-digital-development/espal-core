package assethandler

import (
	"bytes"
	"mime"
	"os/exec"
	"strconv"
	"strings"

	"github.com/juju/errors"
)

const (
	pngQuantMinSpeed = 1
	pngQuantMaxSpeed = 10
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
			shrunkenSizeInBytes, err := h.pngQuant(data, "1")
			if err != nil {
				loopErr = errors.Trace(err)
				return false
			}
			if len(shrunkenSizeInBytes) < len(data) {
				data = shrunkenSizeInBytes
			}
		case "image/jpeg":
			// TODO :: 777777 jpegoptim (wrapper or cmd)
		case "image/gif":
			// TODO :: 777777 gifsicle (wrapper or cmd)
		case "image/svg+xml":
			// TODO :: 777777 svgo (wrapper or cmd)
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

// TODO :: 777777 Move this to it's own image service (along with other optimizers)

func (h *AssetHandler) pngQuant(input []byte, speed string) (output []byte, err error) {
	speedInt, err := strconv.Atoi(speed)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if speedInt < pngQuantMinSpeed || speedInt > pngQuantMaxSpeed {
		return nil, errors.Errorf("speed has to be between %d and %d", pngQuantMinSpeed, pngQuantMaxSpeed)
	}
	cmd := exec.Command("pngquant", "-", "--speed", speed)
	cmd.Stdin = strings.NewReader(string(input))
	var o bytes.Buffer
	cmd.Stdout = &o
	if err := cmd.Run(); err != nil {
		return nil, errors.Trace(err)
	}
	return o.Bytes(), nil
}
