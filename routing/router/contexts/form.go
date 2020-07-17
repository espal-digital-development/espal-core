package contexts

import (
	"mime/multipart"
	"strings"

	"github.com/juju/errors"
)

const multiPartMaxMemory = 1024 * 1024 * 128

// FormContext for form handling.
type FormContext interface {
	MultipartForm(maxMemory int64) (*multipart.Form, error)
	FormFile(string) (multipart.File, *multipart.FileHeader, error)
	FormValue(key string) (string, error)
	FormValues(name string) ([]string, error)
}

// MultipartForm returns the submitted multipart form.
// maxMemory is not the upload size limit, but purely what stays in memory and the remainder being written to the disk.
func (c *HTTPContext) MultipartForm(maxMemory int64) (*multipart.Form, error) {
	if err := c.request.ParseMultipartForm(maxMemory); err != nil {
		return nil, errors.Trace(err)
	}
	return c.request.MultipartForm, nil
}

// FormFile returns the submitted form file by the given key.
func (c *HTTPContext) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.request.FormFile(key)
}

// FormValue returns a submitted form value.
func (c *HTTPContext) FormValue(key string) (string, error) {
	if c.formIsParsed {
		return c.request.Form.Get(key), nil
	}
	if strings.HasPrefix(c.request.Header.Get("Content-Type"), "multipart/form-data") {
		if _, err := c.MultipartForm(multiPartMaxMemory); err != nil {
			return "", errors.Trace(err)
		}
	} else {
		if err := c.request.ParseForm(); err != nil {
			return "", errors.Trace(err)
		}
	}
	c.formIsParsed = true
	return c.request.Form.Get(key), nil
}

// FormValues returns all given values for the given field.
func (c *HTTPContext) FormValues(key string) ([]string, error) {
	if c.formIsParsed {
		return c.request.Form[key], nil

	}
	if strings.HasPrefix(c.request.Header.Get("Content-Type"), "multipart/form-data") {
		if _, err := c.MultipartForm(multiPartMaxMemory); err != nil {
			return nil, errors.Trace(err)
		}
	} else {
		if err := c.request.ParseForm(); err != nil {
			return nil, errors.Trace(err)
		}
	}
	c.formIsParsed = true
	return c.request.Form[key], nil
}
