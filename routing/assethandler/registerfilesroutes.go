package assethandler

import (
	"mime"
	"strings"

	"github.com/juju/errors"
)

func (h *AssetHandler) registerFilesRoutes() error {
	var loopErr error
	h.publicRootFilesStorage.Iterate(func(path string, data []byte, err error) bool {
		if strings.HasPrefix(path, ".") || strings.HasSuffix(path, ".gz") {
			return true
		}
		if err := h.RegisterFileRoute(path, "", data); err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		return true
	})
	if loopErr != nil {
		return errors.Trace(loopErr)
	}

	h.publicFilesStorage.Iterate(func(path string, data []byte, err error) bool {
		if strings.HasPrefix(path, ".") || strings.HasSuffix(path, ".gz") {
			return true
		}
		if err := h.RegisterFileRoute(path, "f/", data); err != nil {
			loopErr = errors.Trace(err)
			return false
		}
		return true
	})
	return errors.Trace(loopErr)
}

// RegisterPublicFileRoute registers a route for the given file path with the default path prefix.
func (h *AssetHandler) RegisterPublicFileRoute(path string, data []byte) error {
	return h.RegisterFileRoute(path, "f/", data)
}

// RegisterFileRoute registers a dynamically created route for the given file path.
func (h *AssetHandler) RegisterFileRoute(path string, prefix string, data []byte) error {
	var gzipData []byte
	var err error
	var allowGzip bool
	// TODO :: 77777 This is a harder problem. On one side you want to have the access to be quick as possible
	// and serve the raw bytes. On the other side you want dynamic interaction with an underlying storage. Yet,
	// this gives overhead. Just giving this route the data and gzipData is not good enough. Also generating
	// compressed variants on-the-fly won't be durable with a lot of files.
	// It needs a new Storage hybrid that can also give compressed variants and has a hybrid of filesystem access
	// but also options to keep certains files up until certain sizes or a total treshhold in memory. This engine
	// would also need to hold metadata of the files in memory, as determining it on each load will be too taxing.
	// if _, err := os.Stat(path + ".gz"); !os.IsNotExist(err) {
	gzipData, err = h.convertToGzip(data)
	if err != nil {
		return errors.Trace(err)
	}
	allowGzip = true
	// }
	prefix = strings.Trim(prefix, "/")
	if len(prefix) == 0 {
		prefix = "/"
	} else {
		prefix = "/" + prefix + "/"
	}
	// TODO :: 7 For very big systems having all the data and gzipData in-memory might be
	//           an issue. Need some smart logic that detects which files are accessed
	//           most and to priotize keeping those in memory.
	return h.routerService.RegisterRoute(prefix+path, &route{
		data:        data,
		gzipData:    gzipData,
		contentType: mime.TypeByExtension(path),
		cacheMaxAge: "",
		allowGzip:   allowGzip,
	})
}

// UnregisterFileRoute unregisters a dynamically created route for the given file path.
func (h *AssetHandler) UnregisterFileRoute(path string) error {
	return h.routerService.UnregisterRoute(path)
}
