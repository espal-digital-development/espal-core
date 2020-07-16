// +build !windows

package assethandler

import (
	"bytes"

	"github.com/andybalholm/brotli"
	"github.com/juju/errors"
)

func (h *AssetHandler) convertToBrotli(bts []byte) ([]byte, error) {
	var b bytes.Buffer
	w := brotli.NewWriterLevel(&b, brotli.BestCompression)
	_, err := w.Write(bts)
	if err != nil {
		return make([]byte, 0), errors.Trace(err)
	}
	if err := w.Close(); err != nil {
		return make([]byte, 0), errors.Trace(err)
	}
	return b.Bytes(), nil
}
