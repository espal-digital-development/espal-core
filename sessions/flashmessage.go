package sessions

// FlashMessageType describes the type the FlashMessage was.
type FlashMessageType byte

// FlashMessageType defines what type of message the FlashMessage was.
const (
	FlashMessageTypeInfo FlashMessageType = iota + 1
	FlashMessageTypeSuccess
	FlashMessageTypeWarning
	FlashMessageTypeError
)

// Message represents an envelope that can be passed through the session to display messages over different pages.
type Message interface {
	Type() FlashMessageType
	SetType(_type FlashMessageType)
	Message() string
	SetMessage(message string)
	IsSet() bool
	Success(message string)
	Warning(message string)
	Error(message string)
	ClassName() string
	Icon() string
}

// FlashMessage are messages that are meant to be consumed after being read and thus deleted.
type FlashMessage struct {
	_type   FlashMessageType
	message string
}

// Type returns the type.
func (m *FlashMessage) Type() FlashMessageType {
	return m._type
}

// SetType sets the type.
func (m *FlashMessage) SetType(_type FlashMessageType) {
	m._type = _type
}

// Message returns the message.
func (m *FlashMessage) Message() string {
	return m.message
}

// SetMessage sets the Message.
func (m *FlashMessage) SetMessage(message string) {
	m.message = message
}

// IsSet returns an indicator boolean wether the FlashMessage was already set.
func (m *FlashMessage) IsSet() bool {
	return m._type != 0
}

// Info sets a message of the type FlashMessageTypeInfo.
func (m *FlashMessage) Info(message string) {
	m._type = FlashMessageTypeInfo
	m.message = message
}

// Success sets a message of the type FlashMessageTypeSuccess.
func (m *FlashMessage) Success(message string) {
	m._type = FlashMessageTypeSuccess
	m.message = message
}

// Warning sets a message of the type FlashMessageTypeWarning.
func (m *FlashMessage) Warning(message string) {
	m._type = FlashMessageTypeWarning
	m.message = message
}

// Error sets a message of the type FlashMessageTypeError.
func (m *FlashMessage) Error(message string) {
	m._type = FlashMessageTypeError
	m.message = message
}

// ClassName returns the CSS class name that suits the message type.
func (m *FlashMessage) ClassName() string {
	switch m._type {
	case FlashMessageTypeInfo:
		return "flash info"
	case FlashMessageTypeSuccess:
		return "flash success"
	case FlashMessageTypeWarning:
		return "flash warning"
	case FlashMessageTypeError:
		return "flash error"
	}
	return ""
}

// Icon returns the icon that suits the message type.
func (m *FlashMessage) Icon() string {
	switch m._type {
	case FlashMessageTypeInfo:
		return "✯ "
	case FlashMessageTypeSuccess:
		return "✓ "
	case FlashMessageTypeWarning:
		return "⚠ "
	case FlashMessageTypeError:
		return "✗ "
	}
	return ""
}

// TODO :: 77 Spawn FlashMessage from a factory too?

// NewFlashMessage returns a new instance of FlashMessage.
func NewFlashMessage() *FlashMessage {
	return &FlashMessage{}
}
