package mailer

import (
	"github.com/go-gomail/gomail"
)

var _ Engine = &Mailer{}

// Engine represents a mailer engine's instructions.
type Engine interface {
	NewMessage() Data
	Send(message Data) error
}

// Mailer sending engine.
type Mailer struct {
	dailer *gomail.Dialer
}

// NewMessage returns a new instance of Data.
func (m *Mailer) NewMessage() Data {
	return &Message{
		headers: map[string][]string{},
		msg:     gomail.NewMessage(),
	}
}

// Send will send the supplied message.
func (m *Mailer) Send(message Data) error {
	msg := gomail.NewMessage()
	msg.SetBody(message.GetContentType(), message.GetBody())
	msg.SetHeaders(message.GetHeaders())
	return m.dailer.DialAndSend(msg)
}

// New returns a new instance of Mailer.
func New(host string, port int, username, password string) *Mailer {
	return &Mailer{
		dailer: gomail.NewDialer(host, port, username, password),
	}
}
