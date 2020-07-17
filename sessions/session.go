package sessions

import (
	"bytes"
	"time"
)

var _ Session = &UserSession{}

// SessionKeys are the keys that can be used for Session index keys.
const (
	SessionKeyFormToken uint8 = iota + 1
	SessionKeyUserID
	SessionKeyRefererURL
	SessionKeyFlashMessageType
	SessionKeyFlashMessageMessage
)

// Session represents an object holding session for a specific object.
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
func (s *UserSession) ID() string {
	return s.id
}

// CreatedAt returns the time the session was created.
func (s *UserSession) CreatedAt() time.Time {
	return s.createdAt
}

// AccessedAt returns the time the session was last accessed.
func (s *UserSession) AccessedAt() time.Time {
	return s.accessedAt
}

// Timeout returns the duration of the session before it will time out since it's last access time.
func (s *UserSession) Timeout() time.Duration {
	return s.timeout
}

// SetTimeout allows the timeout to be modified dynamically.
// It's discouraged to use this other than from within the Sessions.
func (s *UserSession) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

// IsNew returns an indicator if session is new and not saved yet.
func (s *UserSession) IsNew() bool {
	return s.isNew
}

// IsModified returns an indicator if session is has been modified.
func (s *UserSession) IsModified() bool {
	return s.isModified
}

// Set sets a bytes slice value for the given key.
func (s *UserSession) Set(k uint8, v []byte) {
	if _, ok := s.values[k]; ok {
		if bytes.Equal(s.values[k], v) {
			return
		}
	}
	s.isModified = true
	s.accessedAt = time.Now()
	s.values[k] = v
}

// Get returns a bytes slice based on the given key.
func (s *UserSession) Get(k uint8) ([]byte, bool) {
	s.accessedAt = time.Now()
	v, ok := s.values[k]
	return v, ok
}

// Unset unsets a bytes slice value based on given key.
func (s *UserSession) Unset(k uint8) {
	if _, ok := s.values[k]; !ok {
		return
	}
	s.isModified = true
	delete(s.values, k)
}

// HasFlashMessage returns an indicator if a consumable flash message is present in the session.
func (s *UserSession) HasFlashMessage() bool {
	_, ok := s.Get(SessionKeyFlashMessageType)
	return ok
}

// GetFlashMessage gets and consumes the flash message for this context.
func (s *UserSession) GetFlashMessage() Message {
	flashMessage := &FlashMessage{}

	fmType, ok := s.Get(SessionKeyFlashMessageType)
	if !ok {
		return nil
	}
	flashMessage._type = FlashMessageType(fmType[0])
	s.Unset(SessionKeyFlashMessageType)

	fmMessage, ok := s.Get(SessionKeyFlashMessageMessage)
	if !ok {
		return nil
	}
	flashMessage.message = string(fmMessage)
	s.Unset(SessionKeyFlashMessageMessage)

	return flashMessage
}

// SetFlashMessage sets the consumable flash message for this context.
func (s *UserSession) SetFlashMessage(fm Message) {
	s.Set(SessionKeyFlashMessageType, []byte{byte(fm.Type())})
	s.Set(SessionKeyFlashMessageMessage, []byte(fm.Message()))
}

// AllValues returns all byte slice values present in this session.
func (s *UserSession) AllValues() map[uint8][]byte {
	return s.values
}

// NewSession returns a new instance of Session.
func NewSession(id string, createdAt time.Time, accessedAt time.Time, timeout time.Duration,
	uintValues map[uint8]uint, values map[uint8][]byte) *UserSession {
	return &UserSession{
		id:         id,
		createdAt:  createdAt,
		accessedAt: accessedAt,
		timeout:    timeout,
		values:     values,
	}
}
