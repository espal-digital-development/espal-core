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
func (m *Message) SetHeader(field string, value ...string) {
	m.headers[field] = value
	m.msg.SetHeader(field, value...)
}

// GetHeaders returns all the message's set headers.
func (m *Message) GetHeaders() map[string][]string {
	return m.headers
}

// SetBody sets the body of the message. It replaces any content previously set
// by SetBody, AddAlternative or AddAlternativeWriter.
func (m *Message) SetBody(contentType, body string) {
	m.contentType = contentType
	m.body = body
	m.msg.SetBody(contentType, body)
}

// GetContentType returns the message's content type.
func (m *Message) GetContentType() string {
	return m.contentType
}

// GetBody returns the message's body content.
func (m *Message) GetBody() string {
	return m.body
}
