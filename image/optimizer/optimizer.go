package optimizer

import (
	"github.com/espal-digital-development/espal-core/config"
	"github.com/juju/errors"
)

// Optimizable represents an object that can optimize data based on it's type.
type Optimizable interface {
	ForMIMEType(data []byte, mimeType string) (newData []byte, changed bool, err error)
}

// Optimizer provides convertors that can optimize the meta-data and compression and clears unneeded data from given
// data.
type Optimizer struct {
	configService config.Config
}

// ForMIMEType optimizes the data based on the given MIME type.
func (o *Optimizer) ForMIMEType(data []byte, mimeType string) (newData []byte, changed bool, err error) {
	switch mimeType {
	case "image/png":
		if !o.configService.OptimizePngs() {
			return nil, false, nil
		}
		newData, err = o.pngQuant(data, 1)
	case "image/jpeg":
		if !o.configService.OptimizeJpegs() {
			return nil, false, nil
		}
		newData, err = o.jpegOptim(data, 100)
	case "image/gif":
		if !o.configService.OptimizeGifs() {
			return nil, false, nil
		}
		newData, err = o.gifsicle(data, 3, 100)
	case "image/svg+xml":
		if !o.configService.OptimizeSvgs() {
			return nil, false, nil
		}
		newData, err = o.svgo(data)
	default:
		return
	}
	if err != nil {
		err = errors.Trace(err)
		return
	}
	if len(newData) < len(data) {
		return newData, true, nil
	}
	return
}

// New returns a new instance of Optimizer.
func New(configService config.Config) (*Optimizer, error) {
	o := &Optimizer{
		configService: configService,
	}
	return o, nil
}
