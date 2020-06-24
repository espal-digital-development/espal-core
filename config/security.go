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

// SecurityGlobalAuthentication returns if the global authentication is
// enabled. This will cause a global login barrier over the whole
// environment, like BasicAuth, but managed by the normal requests.
func (configuration *Configuration) SecurityGlobalAuthentication() bool {
	return configuration.security.GlobalAuthentication
}

// SecurityBcryptRounds returns the bcrypt encryption rounds.
// The value lies between 1 and 30 and relies on gradual improvment
// of hardware.
func (configuration *Configuration) SecurityBcryptRounds() int {
	return configuration.security.BcryptRounds
}

// SecurityFormTokenLifespan returns the form token's lifespan.
func (configuration *Configuration) SecurityFormTokenLifespan() time.Duration {
	return configuration.security.FormTokenLifespan
}

// SecurityFormTokenCleanupInterval returns the form token's cleanup interval.
func (configuration *Configuration) SecurityFormTokenCleanupInterval() time.Duration {
	return configuration.security.FormTokenCleanupInterval
}
