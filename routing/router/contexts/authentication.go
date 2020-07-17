package contexts

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/sessions"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/juju/errors"
)

// AuthenticationContext for authentication handling and provisioning.
type AuthenticationContext interface {
	SetSessionValue(key uint8, value []byte) error
	GetSessionValue(key uint8) ([]byte, bool, error)
	UnsetSessionValue(key uint8) error
	SaveSessionIfNeeded() error

	IsLoggedIn() bool
	Login(userID string, rememberMe bool) error
	Logout() error
	GetUser() (*user.User, bool, error)

	HasUserRight(string) bool
	HasUserRightOrForbid(userRightName string) bool

	IsDevelopment() bool
	HasAdminAccess() bool
	HasPprofEnabled() bool
	AdminURL() string
	PprofURL() string
}

func (c *HTTPContext) newSession() (sessions.Session, error) {
	session, cookie, err := c.sessionsFactory.New()
	if err != nil {
		return nil, errors.Trace(err)
	}
	c.session = session
	http.SetCookie(c.responseWriter, cookie)
	return session, nil
}

func (c *HTTPContext) getSession() (sessions.Session, bool, error) {
	if c.session != nil {
		return c.session, true, nil
	}
	cookie, err := c.request.Cookie(c.configService.SessionCookieName())
	if err == http.ErrNoCookie {
		return nil, false, nil
	}
	var ok bool
	c.session, ok, err = c.sessionsFactory.Get(cookie.Value)
	return c.session, ok, errors.Trace(err)
}

// SetSessionValue sets the value for the given session key.
func (c *HTTPContext) SetSessionValue(key uint8, value []byte) error {
	session, ok, err := c.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		session, err = c.newSession()
		if err != nil {
			return errors.Trace(err)
		}
	}
	session.Set(key, value)
	return nil
}

// GetSessionValue returns the value for the given key.
func (c *HTTPContext) GetSessionValue(key uint8) ([]byte, bool, error) {
	session, ok, err := c.getSession()
	if err != nil {
		return nil, false, errors.Trace(err)
	}
	if !ok {
		return nil, false, nil
	}
	value, ok := session.Get(key)
	if !ok {
		return nil, false, nil
	}
	return value, true, nil
}

// UnsetSessionValue will remove the value (if it exists) for the given key.
func (c *HTTPContext) UnsetSessionValue(key uint8) error {
	session, ok, err := c.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		return nil
	}
	session.Unset(key)
	return nil
}

// SaveSessionIfNeeded saves the session if it exists or has been touched.
func (c *HTTPContext) SaveSessionIfNeeded() error {
	if c.session == nil {
		return nil
	}
	return c.sessionsFactory.Save(c.session)
}

// IsLoggedIn returns if the user, if present, on the session is logged in.
func (c *HTTPContext) IsLoggedIn() bool {
	user, ok, err := c.GetUser()
	if !ok {
		return false
	}
	if !user.Active() {
		if err := c.Logout(); err != nil {
			c.loggerService.Error(errors.ErrorStack(err))
			return false
		}
		return false
	}
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	return ok
}

// Login will register the User's ID to the session and assign it to the existing cookie. If no cookie is present, a
// new one will be made.
func (c *HTTPContext) Login(userID string, rememberMe bool) error {
	session, ok, err := c.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		session, err = c.newSession()
		if err != nil {
			return errors.Trace(err)
		}
	}
	if rememberMe {
		http.SetCookie(c.responseWriter, c.sessionsFactory.SetRememberMe(session))
	}
	session.Set(sessions.SessionKeyUserID, []byte(userID))
	if c.user != nil && userID != c.user.ID() {
		c.user = nil
	}

	return nil
}

// Logout will remove the User's ID from the session and wipe the cookie if it's the last value.
func (c *HTTPContext) Logout() error {
	if !c.IsLoggedIn() {
		return nil
	}
	session, ok, err := c.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if ok {
		session.Unset(sessions.SessionKeyUserID)
	}
	c.user = nil
	return nil
}

// GetUser returns the logged-in User.
func (c *HTTPContext) GetUser() (*user.User, bool, error) {
	session, ok, err := c.getSession()
	if err != nil || !ok {
		return nil, ok, errors.Trace(err)
	}
	userID, ok := session.Get(sessions.SessionKeyUserID)
	if !ok {
		return nil, false, nil
	}
	c.user, ok, err = c.userStore.GetOneActive(string(userID))
	return c.user, ok, errors.Trace(err)
}

// HasUserRight determines if the logged-in User has the required UserRight.
func (c *HTTPContext) HasUserRight(userRightName string) bool {
	user, ok, err := c.GetUser()
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	if !ok {
		return false
	}
	hasUserRight, err := c.userStore.HasUserRight(user, userRightName)
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	return hasUserRight
}

// HasUserRightOrForbid determines if the logged-in User has the UserRight.
// When it doesn't, it instantly sets the http status to unauthorized.
func (c *HTTPContext) HasUserRightOrForbid(userRightName string) bool {
	if c.HasUserRight(userRightName) {
		return true
	}
	c.RenderUnauthorized()
	return false
}

// HasAdminAccess determines if the logged-in User has access to the admin section of the system.
func (c *HTTPContext) HasAdminAccess() bool {
	return c.HasUserRight("AccessAdminSection")
}

// IsDevelopment returns an indicator if the project is in development mode.
func (c *HTTPContext) IsDevelopment() bool {
	return c.configService.Development()
}

// HasPprofEnabled determines if the logged-in User has access to the pprof section of the system.
func (c *HTTPContext) HasPprofEnabled() bool {
	return c.configService.Pprof()
}

// AdminURL returns the configurated Admin section's URL prefix.
func (c *HTTPContext) AdminURL() string {
	return c.configService.AdminURL()
}

// PprofURL returns the configurated Pprof section's URL prefix.
func (c *HTTPContext) PprofURL() string {
	return c.configService.PprofURL()
}
