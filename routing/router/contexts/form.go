package contexts

import (
	"mime/multipart"

	"github.com/juju/errors"
)

// FormContext for form handling.
type FormContext interface {
	MultipartForm(maxMemory int64) (*multipart.Form, error)
	FormFile(string) (multipart.File, *multipart.FileHeader, error)
	FormValue(key string) (string, error)
	FormValues(name string) ([]string, error)
}

// MultipartForm returns the submitted multipart form.
// maxMemory is not the upload size limit, but purely what stays in memory and
// the remainder being written to the disk.
func (httpContext *HTTPContext) MultipartForm(maxMemory int64) (*multipart.Form, error) {
	if err := httpContext.request.ParseMultipartForm(maxMemory); err != nil {
		return nil, errors.Trace(err)
	}
	return httpContext.request.MultipartForm, nil
}

// FormFile returns the submitted form file by the given key.
func (httpContext *HTTPContext) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return httpContext.request.FormFile(key)
}

// FormValue returns a submitted form value.
func (httpContext *HTTPContext) FormValue(key string) (string, error) {
	if !httpContext.formIsParsed {
		if err := httpContext.request.ParseForm(); err != nil {
			return "", errors.Trace(err)
		}
	}
	httpContext.formIsParsed = true
	return httpContext.request.Form.Get(key), nil
}

// FormValues returns all given values for the given field.
func (httpContext *HTTPContext) FormValues(key string) ([]string, error) {
	if !httpContext.formIsParsed {
		if err := httpContext.request.ParseForm(); err != nil {
			return nil, errors.Trace(err)
		}
	}
	httpContext.formIsParsed = true
	return httpContext.request.Form[key], nil
}
