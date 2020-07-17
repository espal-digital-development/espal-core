package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"

	"github.com/espal-digital-development/espal-core/config"
	store "github.com/espal-digital-development/espal-core/stores/session"
	"github.com/juju/errors"
)

var _ Factory = &Sessions{}

// Factory represents an object that can manage session data and logic.
type Factory interface {
	SetRememberMe(session Session) *http.Cookie
	New() (Session, *http.Cookie, error)
	Exists(sessionID string) (bool, error)
	Get(sessionID string) (Session, bool, error)
	Save(session Session) error
}

// Sessions contains all session logic based on a cookie value.
type Sessions struct {
	configService config.Config
	sessionStore  store.Store
}

func (s *Sessions) generateCookie(session Session, expiration time.Time) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Name = s.configService.SessionCookieName()
	cookie.Value = session.ID()
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.Expires = expiration
	cookie.SameSite = http.SameSiteStrictMode
	return cookie
}

// SetRememberMe will modify the cookie expiration time to work according to the remember me setting.
func (s *Sessions) SetRememberMe(session Session) *http.Cookie {
	session.SetTimeout(s.configService.SessionRememberMeExpiration())
	return s.generateCookie(session, time.Now().Add(s.configService.SessionRememberMeExpiration()))
}

// New generates a new internal session and returns it's instance.
func (s *Sessions) New() (Session, *http.Cookie, error) {
	random := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, random); err != nil {
		return nil, nil, errors.Trace(err)
	}

	session := &UserSession{
		id:         base64.URLEncoding.EncodeToString(random),
		createdAt:  time.Now(),
		accessedAt: time.Now(),
		timeout:    s.configService.SessionExpiration(),
		values:     map[uint8][]byte{},
		isNew:      true,
	}

	cookie := s.generateCookie(session, time.Now().Add(s.configService.SessionExpiration()))

	return session, cookie, nil
}

// Exists returns wether the given ID (still) has an existing Session.
func (s *Sessions) Exists(hash string) (bool, error) {
	return s.sessionStore.HashExists(hash)
}

// Get returns the session based on the given header's cookie.
func (s *Sessions) Get(hash string) (Session, bool, error) {
	sess, ok, err := s.sessionStore.GetOneByHash(hash)
	if err != nil || !ok {
		return nil, ok, errors.Trace(err)
	}

	accessedAt := time.Now()
	if sess.UpdatedAt() != nil {
		accessedAt = *sess.UpdatedAt()
	}

	session := &UserSession{
		id:         sess.Hash(),
		createdAt:  sess.CreatedAt(),
		accessedAt: accessedAt,
		timeout:    sess.Timeout(),
		values:     map[uint8][]byte{},
	}

	data, err := sess.GetDataAsJSON()
	if err != nil {
		return nil, false, errors.Trace(err)
	}
	for k, v := range data.All() {
		session.values[k] = []byte(v)
	}

	return session, true, nil
}

// Save saves the given session to the session's storage engine.
func (s *Sessions) Save(session Session) error {
	// TODO :: 77 AccessedAt is always updating, but way too heavy. Maybe needs a more
	// modern way of approaching? Like only update when a certain amount of seconds or
	// minutes had passed? Or it would expire it or not; but will not work safely either

	// TODO :: 77 Delete when exists, but data ended up empty (unset cookie too?)

	if !session.IsModified() {
		return nil
	}

	// TODO :: 77 This is sloppy and needs a prettier implementation
	data := store.NewDataEntries()
	for key, value := range session.AllValues() {
		data.Set(key, string(value))
	}

	if session.IsNew() {
		if err := s.sessionStore.Create(session.ID(), session.Timeout(), data); err != nil {
			return errors.Trace(err)
		}
	} else {
		if err := s.sessionStore.Update(session.ID(), session.Timeout(), data); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

// New returns a new instance of Sessionsessions.
func New(configService config.Config, sessionStore store.Store) *Sessions {
	return &Sessions{
		configService: configService,
		sessionStore:  sessionStore,
	}
}
