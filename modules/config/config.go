package config

import "github.com/espal-digital-development/espal-core/config"

type Config struct {
	service config.Config
}

// SetService sets the core Config service.
func (c *Config) SetService(service config.Config) {
	c.service = service
}

// GetService returns the core Config service.
func (c *Config) GetService() config.Config {
	return c.service
}

// New returns a new instance of Config.
func New() (*Config, error) {
	c := &Config{}
	return c, nil
}
