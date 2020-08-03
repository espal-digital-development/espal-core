package themes

import (
	"io"

	"github.com/juju/errors"
)

var _ Viewable = &View{}

// Viewable represents an object that provides visual rendering processing.
type Viewable interface {
	Code() string
	Render(w io.Writer, data DataStore) error
}

// View manages visual processing.
type View struct {
	code     string
	callback func(w io.Writer, data DataStore) error
}

// Code returns the unique View code.
func (v *View) Code() string {
	return v.code
}

// Render returns the View's final output.
func (v *View) Render(w io.Writer, data DataStore) error {
	if v.callback == nil {
		return errors.Errorf("no callback set for view with code `%s`", v.code)
	}
	return v.callback(w, data)
}

// SetCallback sets the callback for the Render invocation.
func (v *View) SetCallback(callback func(w io.Writer, data DataStore) error) {
	v.callback = callback
}

// NewView returns a new instance of View.
func NewView(code string) *View {
	return &View{
		code: code,
	}
}
