package config

type Config struct{}

// New returns a new instance of Config.
func New() (*Config, error) {
	c := &Config{}
	return c, nil
}
