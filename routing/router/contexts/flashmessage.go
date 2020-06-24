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

// HasFlashMessage returns an indicator if a consumable flash message
// is present in the session.
func (httpContext *HTTPContext) HasFlashMessage() bool {
	session, ok, err := httpContext.getSession()
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	if !ok {
		return false
	}
	if session == nil {
		return false
	}
	if httpContext.hasFlashMessage {
		return true
	}
	// Buffer so it understands the FlashMessage is shown on this page,
	// even when it was already consumed.
	httpContext.hasFlashMessage = session.HasFlashMessage()
	return httpContext.hasFlashMessage
}

// GetFlashMessage gets and consumes the flash message for this httpContext.
func (httpContext *HTTPContext) GetFlashMessage() sessions.Message {
	if httpContext.flashMessageBuffer != nil {
		return httpContext.flashMessageBuffer
	}
	session, ok, err := httpContext.getSession()
	if err != nil {
		httpContext.loggerService.Error(err.Error())
		return nil
	}
	if !ok {
		return nil
	}
	// Buffer it for re-use in this cycle as it only gets consumed once
	httpContext.flashMessageBuffer = session.GetFlashMessage()
	return httpContext.flashMessageBuffer
}

// setFlashMessage sets the flash type and message for the context.
func (httpContext *HTTPContext) setFlashMessage(messageType sessions.FlashMessageType, message string) error {
	session, ok, err := httpContext.getSession()
	if err != nil {
		return errors.Trace(err)
	}
	if !ok {
		var err error
		session, err = httpContext.newSession()
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
func (httpContext *HTTPContext) SetFlashInfoMessage(message string) error {
	return httpContext.setFlashMessage(sessions.FlashMessageTypeInfo, message)
}

// SetFlashSuccessMessage sets an success message for the context.
func (httpContext *HTTPContext) SetFlashSuccessMessage(message string) error {
	return httpContext.setFlashMessage(sessions.FlashMessageTypeSuccess, message)
}

// SetFlashWarningMessage sets an warning message for the context.
func (httpContext *HTTPContext) SetFlashWarningMessage(message string) error {
	return httpContext.setFlashMessage(sessions.FlashMessageTypeWarning, message)
}

// SetFlashErrorMessage sets an error message for the context.
func (httpContext *HTTPContext) SetFlashErrorMessage(message string) error {
	return httpContext.setFlashMessage(sessions.FlashMessageTypeError, message)
}
