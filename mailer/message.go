package mailer

import (
	"github.com/go-gomail/gomail"
)

var _ Data = &Message{}

// Data represents a mail message object.
type Data interface {
	SetHeader(field string, value ...string)
	SetBody(contentType, body string)
	GetHeaders() map[string][]string
	GetContentType() string
	GetBody() string
}

// Message mail message object.
type Message struct {
	msg         *gomail.Message
	headers     map[string][]string
	contentType string
	body        string
}

// SetHeader sets a value to the given header field.
func (message *Message) SetHeader(field string, value ...string) {
	message.headers[field] = value
	message.msg.SetHeader(field, value...)
}

// GetHeaders returns all the message's set headers.
func (message *Message) GetHeaders() map[string][]string {
	return message.headers
}

// SetBody sets the body of the message. It replaces any content previously set
// by SetBody, AddAlternative or AddAlternativeWriter.
func (message *Message) SetBody(contentType, body string) {
	message.contentType = contentType
	message.body = body
	message.msg.SetBody(contentType, body)
}

// GetContentType returns the message's content type.
func (message *Message) GetContentType() string {
	return message.contentType
}

// GetBody returns the message's body content.
func (message *Message) GetBody() string {
	return message.body
}
