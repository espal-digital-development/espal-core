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

	HasAdminAccess() bool
	HasPprofEnabled() bool
	AdminURL() string
	PprofURL() string
}

func (httpContext *HTTPContext) newSession() (sessions.Session, error) {
	session, cookie, err := httpContext.sessionsFactory.New()
	if err != nil {
		return nil, errors.Trace(err)
	}
	httpContext.session = session
	http.SetCookie(httpContext.responseWriter, cookie)
	return session, nil
}

func (httpContext *HTTPContext) getSession() (sessions.Session, bool, error) {
	if httpContext.session != nil {
		return httpContext.session, true, nil
	}
	cookie, err := httpContext.request.Cookie(httpContext.configService.SessionCookieName())
	if err == http.ErrNoCookie {
		return nil, false, nil
	}
	var ok bool
	httpContext.session, ok, err = httpContext.sessionsFactory.Get(cookie.Value)
	return httpContext.session, ok, errors.Trace(err)
}

// SetSessionValue sets the value for the given session key.
func (httpContext *HTTPContext) SetSessionValue(key uint8, value []byte) error {
	session, ok, err := httpContext.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		session, err = httpContext.newSession()
		if err != nil {
			return errors.Trace(err)
		}
	}
	session.Set(key, value)
	return nil
}

// GetSessionValue returns the value for the given key.
func (httpContext *HTTPContext) GetSessionValue(key uint8) ([]byte, bool, error) {
	session, ok, err := httpContext.getSession()
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
func (httpContext *HTTPContext) UnsetSessionValue(key uint8) error {
	session, ok, err := httpContext.getSession()
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
func (httpContext *HTTPContext) SaveSessionIfNeeded() error {
	if httpContext.session == nil {
		return nil
	}
	return httpContext.sessionsFactory.Save(httpContext.session)
}

// IsLoggedIn returns if the user, if present, on the session is logged in.
func (httpContext *HTTPContext) IsLoggedIn() bool {
	user, ok, err := httpContext.GetUser()
	if !ok {
		return false
	}
	if !user.Active() {
		if err := httpContext.Logout(); err != nil {
			httpContext.loggerService.Error(errors.ErrorStack(err))
			return false
		}
		return false
	}
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	return ok
}

// Login will register the User's ID to the session and assign it to the
// existing cookie. If no cookie is present, a new one will be made.
func (httpContext *HTTPContext) Login(userID string, rememberMe bool) error {
	session, ok, err := httpContext.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		session, err = httpContext.newSession()
		if err != nil {
			return errors.Trace(err)
		}
	}
	if rememberMe {
		http.SetCookie(httpContext.responseWriter, httpContext.sessionsFactory.SetRememberMe(session))
	}
	session.Set(sessions.SessionKeyUserID, []byte(userID))
	if httpContext.user != nil && userID != httpContext.user.ID() {
		httpContext.user = nil
	}

	return nil
}

// Logout will remove the User's ID from the session and wipe
// the cookie if it's the last value.
func (httpContext *HTTPContext) Logout() error {
	if !httpContext.IsLoggedIn() {
		return nil
	}
	session, ok, err := httpContext.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if ok {
		session.Unset(sessions.SessionKeyUserID)
	}
	httpContext.user = nil
	return nil
}

// GetUser returns the logged-in User.
func (httpContext *HTTPContext) GetUser() (*user.User, bool, error) {
	session, ok, err := httpContext.getSession()
	if err != nil || !ok {
		return nil, ok, errors.Trace(err)
	}
	userID, ok := session.Get(sessions.SessionKeyUserID)
	if !ok {
		return nil, false, nil
	}
	httpContext.user, ok, err = httpContext.userStore.GetOneActive(string(userID))
	return httpContext.user, ok, errors.Trace(err)
}

// HasUserRight determines if the logged-in User has the required UserRight.
func (httpContext *HTTPContext) HasUserRight(userRightName string) bool {
	user, ok, err := httpContext.GetUser()
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	if !ok {
		return false
	}
	hasUserRight, err := httpContext.userStore.HasUserRight(user, userRightName)
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	return hasUserRight
}

// HasUserRightOrForbid determines if the logged-in User has the UserRight.
// When it doesn't, it instantly sets the http status to unauthorized.
func (httpContext *HTTPContext) HasUserRightOrForbid(userRightName string) bool {
	if httpContext.HasUserRight(userRightName) {
		return true
	}
	httpContext.RenderUnauthorized()
	return false
}

// HasAdminAccess determines if the logged-in User has access to the admin section of the system.
func (httpContext *HTTPContext) HasAdminAccess() bool {
	return httpContext.HasUserRight("AccessAdminSection")
}

// HasPprofEnabled determines if the logged-in User has access to the pprof section of the system.
func (httpContext *HTTPContext) HasPprofEnabled() bool {
	return httpContext.configService.Pprof()
}

// AdminURL returns the configurated Admin section's URL prefix.
func (httpContext *HTTPContext) AdminURL() string {
	return httpContext.configService.AdminURL()
}

// PprofURL returns the configurated Pprof section's URL prefix.
func (httpContext *HTTPContext) PprofURL() string {
	return httpContext.configService.PprofURL()
}
