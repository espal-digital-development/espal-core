package assethandler

import (
	"bytes"
	"compress/gzip"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/routing/router"
	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
)

var _ Handler = &AssetHandler{}

// Handler represents an object that can register and handle routes.
type Handler interface {
	RegisterAll() error
	RegisterPublicFileRoute(path string, data []byte) error
	RegisterFileRoute(path string, prefix string, data []byte) error
	UnregisterFileRoute(path string) error
}

// AssetHandler asset route registration service.
type AssetHandler struct {
	configService          config.Config
	routerService          router.Router
	minifier               *minify.M
	stylesheetsStorage     storage.Storage
	javaScriptStorage      storage.Storage
	imagesStorage          storage.Storage
	publicFilesStorage     storage.Storage
	publicRootFilesStorage storage.Storage
}

// RegisterAll registers all assets as a route equivalent.
func (h *AssetHandler) RegisterAll() error {
	h.minifier.AddFunc("text/css", css.Minify)
	h.minifier.AddFunc("application/javascript", js.Minify)

	// TODO :: 7 Images that get saved in Forms could be optimized and then compressed to gzip/brotli besides the original file and instantly served by keeping a buffer in cache of which files have gzip/brotli variations on disk
	// TODO :: 7 Max-Age header for files too. A bit more complex tho

	if err := h.registerStylesheetsRoutes(); err != nil {
		return errors.Trace(err)
	}
	if err := h.registerJavaScriptsRoutes(); err != nil {
		return errors.Trace(err)
	}
	if err := h.registerImagesRoutes(); err != nil {
		return errors.Trace(err)
	}
	if err := h.registerFilesRoutes(); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (h *AssetHandler) convertToGzip(bts []byte) ([]byte, error) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return make([]byte, 0), errors.Trace(err)
	}
	_, err = w.Write(bts)
	if err != nil {
		return make([]byte, 0), errors.Trace(err)
	}
	if err := w.Close(); err != nil {
		return make([]byte, 0), errors.Trace(err)
	}
	return b.Bytes(), nil
}

// New returns a new instance of Handler.
func New(configService config.Config, routerService router.Router, minifier *minify.M, stylesheetsStorage storage.Storage, javaScriptStorage storage.Storage, imagesStorage storage.Storage, publicFilesStorage storage.Storage, publicRootFilesStorage storage.Storage) *AssetHandler {
	return &AssetHandler{
		configService:          configService,
		routerService:          routerService,
		minifier:               minifier,
		stylesheetsStorage:     stylesheetsStorage,
		javaScriptStorage:      javaScriptStorage,
		imagesStorage:          imagesStorage,
		publicFilesStorage:     publicFilesStorage,
		publicRootFilesStorage: publicRootFilesStorage,
	}
}
