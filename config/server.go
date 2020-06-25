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
func (c *Configuration) ServerHost() string {
	return c.server.Host
}

// ServerPort returns the server port.
func (c *Configuration) ServerPort() int {
	return c.server.Port
}

// ServerHTTPRedirectPort returns the server HTTP redirect port.
func (c *Configuration) ServerHTTPRedirectPort() int {
	return c.server.HTTPRedirectPort
}
