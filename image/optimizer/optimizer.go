package optimizer

import (
	"github.com/juju/errors"
)

// Optimizable represents an object that can optimize
// data based on it's type.
type Optimizable interface {
	ForMIMEType(data []byte, mimeType string) (newData []byte, changed bool, err error)
}

// Optimizer provides convertors that can optimize the meta-data and
// compression and clears unneeded data from given data.
type Optimizer struct {
}

// ForMIMEType optimizes the data based on the given MIME type.
func (o *Optimizer) ForMIMEType(data []byte, mimeType string) (newData []byte, changed bool, err error) {
	switch mimeType {
	case "image/png":
		newData, err = o.pngQuant(data, 1)
	case "image/jpeg":
		newData, err = o.jpegOptim(data, 100)
	case "image/gif":
		newData, err = o.gifsicle(data, 3, 100)
	case "image/svg+xml":
		newData, err = o.svgo(data)
		return
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
func New() (*Optimizer, error) {
	o := &Optimizer{}
	return o, nil
}
