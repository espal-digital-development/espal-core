package config

// Server config entry.
type Server interface {
	ServerHost() string
	ServerPort() int
	ServerHTTPRedirectPort() int
}

type server struct {
	Host             string
	Port             int
	HTTPRedirectPort int `yaml:"httpRedirectPort"`
}

// ServerHost returns the server host.
func (configuration *Configuration) ServerHost() string {
	return configuration.server.Host
}

// ServerPort returns the server port.
func (configuration *Configuration) ServerPort() int {
	return configuration.server.Port
}

// ServerHTTPRedirectPort returns the server HTTP redirect port.
func (configuration *Configuration) ServerHTTPRedirectPort() int {
	return configuration.server.HTTPRedirectPort
}
