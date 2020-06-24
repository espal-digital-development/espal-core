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

// Message represents an envelope that can be passed through
// the session to display messages over different pages.
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

// FlashMessage are messages that are meant to be consumed
// after being read and thus deleted.
type FlashMessage struct {
	_type   FlashMessageType
	message string
}

// Type returns the type.
func (flashMessage *FlashMessage) Type() FlashMessageType {
	return flashMessage._type
}

// SetType sets the type.
func (flashMessage *FlashMessage) SetType(_type FlashMessageType) {
	flashMessage._type = _type
}

// Message returns the message.
func (flashMessage *FlashMessage) Message() string {
	return flashMessage.message
}

// SetMessage sets the Message.
func (flashMessage *FlashMessage) SetMessage(message string) {
	flashMessage.message = message
}

// IsSet returns an indicator boolean wether the FlashMessage was already set.
func (flashMessage *FlashMessage) IsSet() bool {
	return flashMessage._type != 0
}

// Info sets a message of the type FlashMessageTypeInfo.
func (flashMessage *FlashMessage) Info(message string) {
	flashMessage._type = FlashMessageTypeInfo
	flashMessage.message = message
}

// Success sets a message of the type FlashMessageTypeSuccess.
func (flashMessage *FlashMessage) Success(message string) {
	flashMessage._type = FlashMessageTypeSuccess
	flashMessage.message = message
}

// Warning sets a message of the type FlashMessageTypeWarning.
func (flashMessage *FlashMessage) Warning(message string) {
	flashMessage._type = FlashMessageTypeWarning
	flashMessage.message = message
}

// Error sets a message of the type FlashMessageTypeError.
func (flashMessage *FlashMessage) Error(message string) {
	flashMessage._type = FlashMessageTypeError
	flashMessage.message = message
}

// ClassName returns the CSS class name that suits the message type.
func (flashMessage *FlashMessage) ClassName() string {
	switch flashMessage._type {
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
func (flashMessage *FlashMessage) Icon() string {
	switch flashMessage._type {
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
