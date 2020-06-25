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
func (c *Configuration) SessionCookieName() string {
	return c.session.CookieName
}

// SessionExpiration returns the cookie expiration.
func (c *Configuration) SessionExpiration() time.Duration {
	return c.session.Expiration
}

// SessionRememberMeExpiration returns the cookie `remember me` expiration.
func (c *Configuration) SessionRememberMeExpiration() time.Duration {
	return c.session.RememberMeExpiration
}
