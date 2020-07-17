package config

// General config section.
type General interface {
	Development() bool
	Logging() bool
	Pprof() bool
	Languages() []string
	DefaultLanguage() string
}

type general struct {
	Development     bool
	Logging         bool
	Pprof           bool
	Languages       []string `yaml:",flow"`
	DefaultLanguage string   `yaml:"defaultLanguage"`
}

// Development returns the indicator if the application is in development mode and might sacrifice speed over being
// more verbose.
func (c *Configuration) Development() bool {
	return c.general.Development
}

// Logging returns an indicator if the logger should be called when routes are visited, and when info-, warning- and
// errors messages are being reported.
func (c *Configuration) Logging() bool {
	return c.general.Logging
}

// Pprof returns an indicator if pprof logging should be possible and the routes will be made available for the pprof
// tool.
func (c *Configuration) Pprof() bool {
	return c.general.Pprof
}

// Languages returns the list of available language codes.
func (c *Configuration) Languages() []string {
	return c.general.Languages
}

// DefaultLanguage returns the default language code.
func (c *Configuration) DefaultLanguage() string {
	return c.general.DefaultLanguage
}
