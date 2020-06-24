package contexts

import (
	"io"
	"net/http"
	"strings"
)

// RequestContext for request handling.
type RequestContext interface {
	io.Writer
	WriteString(p string) (n int, err error)
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
func (httpContext *HTTPContext) Write(p []byte) (int, error) {
	return httpContext.responseWriter.Write(p)
}

// WriteString writes the given string directly to the response writer.
func (httpContext *HTTPContext) WriteString(p string) (int, error) {
	return httpContext.Write([]byte(p))
}

// RequestBody returns the internal body ReadCloser.
func (httpContext *HTTPContext) RequestBody() io.ReadCloser {
	return httpContext.request.Body
}

// Method returns the request HTTP method.
func (httpContext *HTTPContext) Method() string {
	return httpContext.request.Method
}

// Redirect redirects to the given path in combination with the given HTTP status code.
func (httpContext *HTTPContext) Redirect(path string, code int) {
	http.Redirect(httpContext.responseWriter, httpContext.request, path, code)
}

// GetHeader gets a header value on the HTTP request.
func (httpContext *HTTPContext) GetHeader(key string) string {
	return httpContext.request.Header.Get(key)
}

// SetHeader sets a header value on the HTTP response.
func (httpContext *HTTPContext) SetHeader(key string, value string) {
	httpContext.responseWriter.Header().Set(key, value)
}

// StatusCode returns the status code on the HTTP response.
func (httpContext *HTTPContext) StatusCode() int {
	return httpContext.httpStatusCode
}

// SetStatusCode will set the status code on the HTTP response.
func (httpContext *HTTPContext) SetStatusCode(statusCode int) {
	httpContext.httpStatusCode = statusCode
	httpContext.responseWriter.WriteHeader(statusCode)
}

// SetContentType will set the content type on the HTTP response.
func (httpContext *HTTPContext) SetContentType(contentType string) {
	httpContext.responseWriter.Header().Set("Content-Type", contentType)
}

// AcceptsEncoding returns an indicator if the requesting client accepts the
// given encoding.
func (httpContext *HTTPContext) AcceptsEncoding(encoding string) bool {
	return strings.Contains(httpContext.request.Header.Get("Accept-Encoding"), encoding)
}

// RequestURI returns the unmodified request URI.
func (httpContext *HTTPContext) RequestURI() string {
	return httpContext.request.RequestURI
}

// QueryString returns the raw URL query string.
func (httpContext *HTTPContext) QueryString() string {
	return httpContext.request.URL.RawQuery
}

// QueryValue returns the query string belonging to the given key.
func (httpContext *HTTPContext) QueryValue(key string) string {
	return httpContext.request.URL.Query().Get(key)
}

// Host returns the host part of the URL.
func (httpContext *HTTPContext) Host() string {
	return httpContext.request.Host
}

// Path returns the relative URL path.
func (httpContext *HTTPContext) Path() string {
	return httpContext.request.URL.Path
}

// Referer returns the path this page was referred from before a redirect.
func (httpContext *HTTPContext) Referer() string {
	return httpContext.request.Referer()
}
