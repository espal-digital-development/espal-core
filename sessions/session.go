package sessions

import (
	"bytes"
	"time"
)

var _ Session = &UserSession{}

// SessionKeys are the keys that can be used for Session index keys.
const (
	SessionKeyUserID uint8 = iota + 1
	SessionKeyRefererURL
	SessionKeyFlashMessageType
	SessionKeyFlashMessageMessage
)

// Session represents an object holding session for
// a specific object.
type Session interface {
	ID() string
	CreatedAt() time.Time
	AccessedAt() time.Time
	Timeout() time.Duration
	SetTimeout(timeout time.Duration)
	IsNew() bool
	IsModified() bool
	Set(k uint8, v []byte)
	Get(k uint8) ([]byte, bool)
	Unset(k uint8)
	HasFlashMessage() bool
	SetFlashMessage(fm Message)
	GetFlashMessage() Message
	AllValues() map[uint8][]byte
}

// UserSession instance for a specific visitor.
type UserSession struct {
	id         string
	createdAt  time.Time
	accessedAt time.Time
	timeout    time.Duration
	values     map[uint8][]byte
	isNew      bool
	isModified bool
}

// ID returns the unique ID of this session instance.
func (userSession *UserSession) ID() string {
	return userSession.id
}

// CreatedAt returns the time the session was created.
func (userSession *UserSession) CreatedAt() time.Time {
	return userSession.createdAt
}

// AccessedAt returns the time the session was last accessed.
func (userSession *UserSession) AccessedAt() time.Time {
	return userSession.accessedAt
}

// Timeout returns the duration of the session before it will time out
// since it's last access time.
func (userSession *UserSession) Timeout() time.Duration {
	return userSession.timeout
}

// SetTimeout allows the timeout to be modified dynamically.
// It's discouraged to use this other than from within the Sessions
func (userSession *UserSession) SetTimeout(timeout time.Duration) {
	userSession.timeout = timeout
}

// IsNew returns an indicator if session is new and not saved yet.
func (userSession *UserSession) IsNew() bool {
	return userSession.isNew
}

// IsModified returns an indicator if session is has been modified.
func (userSession *UserSession) IsModified() bool {
	return userSession.isModified
}

// Set sets a bytes slice value for the given key.
func (userSession *UserSession) Set(k uint8, v []byte) {
	if _, ok := userSession.values[k]; ok {
		if bytes.Equal(userSession.values[k], v) {
			return
		}
	}
	userSession.isModified = true
	userSession.accessedAt = time.Now()
	userSession.values[k] = v
}

// Get returns a bytes slice based on the given key.
func (userSession *UserSession) Get(k uint8) ([]byte, bool) {
	userSession.accessedAt = time.Now()
	v, ok := userSession.values[k]
	return v, ok
}

// Unset unsets a bytes slice value based on given key.
func (userSession *UserSession) Unset(k uint8) {
	if _, ok := userSession.values[k]; !ok {
		return
	}
	userSession.isModified = true
	delete(userSession.values, k)
}

// HasFlashMessage returns an indicator if a consumable flash message
// is present in the session.
func (userSession *UserSession) HasFlashMessage() bool {
	_, ok := userSession.Get(SessionKeyFlashMessageType)
	return ok
}

// GetFlashMessage gets and consumes the flash message for this context.
func (userSession *UserSession) GetFlashMessage() Message {
	flashMessage := &FlashMessage{}

	fmType, ok := userSession.Get(SessionKeyFlashMessageType)
	if !ok {
		return nil
	}
	flashMessage._type = FlashMessageType(fmType[0])
	userSession.Unset(SessionKeyFlashMessageType)

	fmMessage, ok := userSession.Get(SessionKeyFlashMessageMessage)
	if !ok {
		return nil
	}
	flashMessage.message = string(fmMessage)
	userSession.Unset(SessionKeyFlashMessageMessage)

	return flashMessage
}

// SetFlashMessage sets the consumable flash message for this context.
func (userSession *UserSession) SetFlashMessage(fm Message) {
	userSession.Set(SessionKeyFlashMessageType, []byte{byte(fm.Type())})
	userSession.Set(SessionKeyFlashMessageMessage, []byte(fm.Message()))
}

// AllValues returns all byte slice values present in this session.
func (userSession *UserSession) AllValues() map[uint8][]byte {
	return userSession.values
}

// NewSession returns a new instance of Session.
func NewSession(id string, createdAt time.Time, accessedAt time.Time, timeout time.Duration, uintValues map[uint8]uint, values map[uint8][]byte) *UserSession {
	return &UserSession{
		id:         id,
		createdAt:  createdAt,
		accessedAt: accessedAt,
		timeout:    timeout,
		values:     values,
	}
}
