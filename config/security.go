package config

import (
	"time"
)

// Security config entry.
type Security interface {
	SecurityGlobalAuthentication() bool
	SecurityBcryptRounds() int
	SecurityFormTokenLifespan() time.Duration
	SecurityFormTokenCleanupInterval() time.Duration
}

type security struct {
	GlobalAuthentication     bool          `yaml:"globalAuthentication"`
	BcryptRounds             int           `yaml:"bcryptRounds"`
	FormTokenLifespan        time.Duration `yaml:"formTokenLifespan"`
	FormTokenCleanupInterval time.Duration `yaml:"formTokenCleanupInterval"`
}

// SecurityGlobalAuthentication returns if the global authentication is enabled. This will cause a global login barrier
// over the whole environment, like BasicAuth, but managed by the normal requests.
func (c *Configuration) SecurityGlobalAuthentication() bool {
	return c.security.GlobalAuthentication
}

// SecurityBcryptRounds returns the bcrypt encryption rounds.
// The value lies between 1 and 30 and relies on gradual improvment of hardware.
func (c *Configuration) SecurityBcryptRounds() int {
	return c.security.BcryptRounds
}

// SecurityFormTokenLifespan returns the form token's lifespan.
func (c *Configuration) SecurityFormTokenLifespan() time.Duration {
	return c.security.FormTokenLifespan
}

// SecurityFormTokenCleanupInterval returns the form token's cleanup interval.
func (c *Configuration) SecurityFormTokenCleanupInterval() time.Duration {
	return c.security.FormTokenCleanupInterval
}
