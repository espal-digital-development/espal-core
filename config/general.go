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

// Development returns the indicator if the application is in
// development mode and might sacrifice speed over being more verbose.
func (configuration *Configuration) Development() bool {
	return configuration.general.Development
}

// Logging returns an indicator if the logger should be called
// when routes are visited, and when info-, warning- and errors
// messages are being reported.
func (configuration *Configuration) Logging() bool {
	return configuration.general.Logging
}

// Pprof returns an indicator if pprof logging should be possible
// and the routes will be made available for the pprof tool.
func (configuration *Configuration) Pprof() bool {
	return configuration.general.Pprof
}

// Languages returns the list of available language codes.
func (configuration *Configuration) Languages() []string {
	return configuration.general.Languages
}

// DefaultLanguage returns the default language code.
func (configuration *Configuration) DefaultLanguage() string {
	return configuration.general.DefaultLanguage
}
