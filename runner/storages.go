package runner

import (
	"github.com/espal-digital-development/espal-core/storage"
)

type storages struct {
	// core only functions as a facility to store base data and make
	// the whole core system independent on the engine
	core                  storage.Storage
	translations          storage.Storage
	assetsPrivateFiles    storage.Storage
	assetsPublicFiles     storage.Storage
	assetsPublicRootFiles storage.Storage
	assetsImages          storage.Storage
	assetsStylesheets     storage.Storage
	assetsJavaScript      storage.Storage
}
