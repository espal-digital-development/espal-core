package config

import (
	"time"
)

// Session config entry.
type Session interface {
	SessionCookieName() string
	SessionExpiration() time.Duration
	SessionRememberMeExpiration() time.Duration
}

type session struct {
	CookieName           string        `yaml:"cookieName"`
	Expiration           time.Duration `yaml:"expiration"`
	RememberMeExpiration time.Duration `yaml:"rememberMeExpiration"`
}

// SessionCookieName returns the cookie name.
func (configuration *Configuration) SessionCookieName() string {
	return configuration.session.CookieName
}

// SessionExpiration returns the cookie expiration.
func (configuration *Configuration) SessionExpiration() time.Duration {
	return configuration.session.Expiration
}

// SessionRememberMeExpiration returns the cookie `remember me` expiration.
func (configuration *Configuration) SessionRememberMeExpiration() time.Duration {
	return configuration.session.RememberMeExpiration
}
