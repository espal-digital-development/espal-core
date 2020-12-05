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
	SecurityJWTSigningMethod() string
	SecurityJWTPassword() string
	SecurityHTTPReferrerPolicy() string
	SecurityHTTPContentSecurityPolicy() string
}

type security struct {
	GlobalAuthentication      bool          `yaml:"globalAuthentication"`
	BcryptRounds              int           `yaml:"bcryptRounds"`
	FormTokenLifespan         time.Duration `yaml:"formTokenLifespan"`
	FormTokenCleanupInterval  time.Duration `yaml:"formTokenCleanupInterval"`
	JWTSigningMethod          string        `yaml:"jwtSigningMethod"`
	JWTPassword               string        `yaml:"jwtPassword"`
	HTTPReferrerPolicy        string        `yaml:"httpReferrerPolicy"`
	HTTPContentSecurityPolicy string        `yaml:"httpContentSecurityPolicy"`
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

// SecurityJWTSigningMethod returns the JWT Signing Method.
func (c *Configuration) SecurityJWTSigningMethod() string {
	return c.security.JWTSigningMethod
}

// SecurityJWTPassword returns the JWT Password.
func (c *Configuration) SecurityJWTPassword() string {
	return c.security.JWTPassword
}

// SecurityHTTPReferrerPolicy returns the HTTP ReferrerPolicy.
func (c *Configuration) SecurityHTTPReferrerPolicy() string {
	return c.security.HTTPReferrerPolicy
}

// SecurityHTTPContentSecurityPolicy returns the HTTP ContentSecurityPolicy.
func (c *Configuration) SecurityHTTPContentSecurityPolicy() string {
	return c.security.HTTPContentSecurityPolicy
}
