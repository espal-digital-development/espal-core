// +build windows

package assethandler

var errBrotliNotSupportedOnWindows = errors.New("brotli is not supported on Windows.you can disable it in the config.yml"

func (h *AssetHandler) convertToBrotli(bts []byte) ([]byte, error) {
	return nil, errors.Trace(errBrotliNotSupportedOnWindows)
}
