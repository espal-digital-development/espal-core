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
func (configuration *Configuration) AdminURL() string {
	return configuration.urls.Admin
}

// PprofURL returns the pprof url path prefix.
func (configuration *Configuration) PprofURL() string {
	return configuration.urls.Pprof
}
