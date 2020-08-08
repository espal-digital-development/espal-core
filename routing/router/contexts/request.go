package contexts

import (
	"io"
	"net/http"
	"strings"
)

// RequestContext for request handling.
type RequestContext interface {
	io.Writer
	WriteString(p string)
	RequestBody() io.ReadCloser
	Method() string
	Redirect(path string, statusCode int)
	GetHeader(key string) string
	SetHeader(key string, value string)
	StatusCode() int
	SetStatusCode(int)
	SetContentType(string)
	AcceptsEncoding(encoding string) bool
	RequestURI() string
	QueryString() string
	QueryValue(string) string
	Host() string
	Path() string
	Referer() string
}

// Write writes the given bytes slice directly to the response writer.
func (c *HTTPContext) Write(p []byte) (int, error) {
	return c.responseWriter.Write(p)
}

// WriteString writes the given string directly to the response writer.
func (c *HTTPContext) WriteString(p string) {
	c.Write([]byte(p))
}

// RequestBody returns the internal body ReadCloser.
func (c *HTTPContext) RequestBody() io.ReadCloser {
	return c.request.Body
}

// Method returns the request HTTP method.
func (c *HTTPContext) Method() string {
	return c.request.Method
}

// Redirect redirects to the given path in combination with the given HTTP status code.
func (c *HTTPContext) Redirect(path string, code int) {
	http.Redirect(c.responseWriter, c.request, path, code)
}

// GetHeader gets a header value on the HTTP request.
func (c *HTTPContext) GetHeader(key string) string {
	return c.request.Header.Get(key)
}

// SetHeader sets a header value on the HTTP response.
func (c *HTTPContext) SetHeader(key string, value string) {
	c.responseWriter.Header().Set(key, value)
}

// StatusCode returns the status code on the HTTP response.
func (c *HTTPContext) StatusCode() int {
	return c.httpStatusCode
}

// SetStatusCode will set the status code on the HTTP response.
func (c *HTTPContext) SetStatusCode(statusCode int) {
	c.httpStatusCode = statusCode
	c.responseWriter.WriteHeader(statusCode)
}

// SetContentType will set the content type on the HTTP response.
func (c *HTTPContext) SetContentType(contentType string) {
	c.responseWriter.Header().Set("Content-Type", contentType)
}

// AcceptsEncoding returns an indicator if the requesting client accepts the given encoding.
func (c *HTTPContext) AcceptsEncoding(encoding string) bool {
	return strings.Contains(c.request.Header.Get("Accept-Encoding"), encoding)
}

// RequestURI returns the unmodified request URI.
func (c *HTTPContext) RequestURI() string {
	return c.request.RequestURI
}

// QueryString returns the raw URL query string.
func (c *HTTPContext) QueryString() string {
	return c.request.URL.RawQuery
}

// QueryValue returns the query string belonging to the given key.
func (c *HTTPContext) QueryValue(key string) string {
	return c.request.URL.Query().Get(key)
}

// Host returns the host part of the URL.
func (c *HTTPContext) Host() string {
	return c.request.Host
}

// Path returns the relative URL path.
func (c *HTTPContext) Path() string {
	return c.request.URL.Path
}

// Referer returns the path this page was referred from before a redirect.
func (c *HTTPContext) Referer() string {
	return c.request.Referer()
}
