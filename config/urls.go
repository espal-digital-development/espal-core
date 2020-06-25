package config

// URLs config section.
type URLs interface {
	AdminURL() string
	PprofURL() string
}

type urls struct {
	Admin string
	Pprof string
}

// nolint:stylecheck
// AdminURL returns the admin url path prefix.
func (c *Configuration) AdminURL() string {
	return c.urls.Admin
}

// PprofURL returns the pprof url path prefix.
func (c *Configuration) PprofURL() string {
	return c.urls.Pprof
}
