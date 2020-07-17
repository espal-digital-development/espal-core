package contexts

import (
	"github.com/espal-digital-development/espal-core/sessions"
	"github.com/juju/errors"
)

// FlashMessageContext for flash message handling.
type FlashMessageContext interface {
	HasFlashMessage() bool
	GetFlashMessage() sessions.Message
	SetFlashInfoMessage(string) error
	SetFlashSuccessMessage(string) error
	SetFlashWarningMessage(string) error
	SetFlashErrorMessage(string) error
}

// HasFlashMessage returns an indicator if a consumable flash message is present in the session.
func (c *HTTPContext) HasFlashMessage() bool {
	session, ok, err := c.getSession()
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	if !ok {
		return false
	}
	if session == nil {
		return false
	}
	if c.hasFlashMessage {
		return true
	}
	// Buffer so it understands the FlashMessage is shown on this page,
	// even when it was already consumed.
	c.hasFlashMessage = session.HasFlashMessage()
	return c.hasFlashMessage
}

// GetFlashMessage gets and consumes the flash message for this httpContext.
func (c *HTTPContext) GetFlashMessage() sessions.Message {
	if c.flashMessageBuffer != nil {
		return c.flashMessageBuffer
	}
	session, ok, err := c.getSession()
	if err != nil {
		c.loggerService.Error(err.Error())
		return nil
	}
	if !ok {
		return nil
	}
	// Buffer it for re-use in this cycle as it only gets consumed once
	c.flashMessageBuffer = session.GetFlashMessage()
	return c.flashMessageBuffer
}

// setFlashMessage sets the flash type and message for the context.
func (c *HTTPContext) setFlashMessage(messageType sessions.FlashMessageType, message string) error {
	session, ok, err := c.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		var err error
		session, err = c.newSession()
		if err != nil {
			return errors.Trace(err)
		}
	}
	flashMessage := sessions.NewFlashMessage()
	flashMessage.SetType(messageType)
	flashMessage.SetMessage(message)
	session.SetFlashMessage(flashMessage)
	return nil
}

// SetFlashInfoMessage sets an info message for the context.
func (c *HTTPContext) SetFlashInfoMessage(message string) error {
	return c.setFlashMessage(sessions.FlashMessageTypeInfo, message)
}

// SetFlashSuccessMessage sets an success message for the context.
func (c *HTTPContext) SetFlashSuccessMessage(message string) error {
	return c.setFlashMessage(sessions.FlashMessageTypeSuccess, message)
}

// SetFlashWarningMessage sets an warning message for the context.
func (c *HTTPContext) SetFlashWarningMessage(message string) error {
	return c.setFlashMessage(sessions.FlashMessageTypeWarning, message)
}

// SetFlashErrorMessage sets an error message for the context.
func (c *HTTPContext) SetFlashErrorMessage(message string) error {
	return c.setFlashMessage(sessions.FlashMessageTypeError, message)
}
