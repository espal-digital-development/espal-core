package login

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

var _ Form = &Login{}

type translator interface {
	Translate(string) string
}

type routeCtx interface {
	contexts.FormContext
	contexts.RequestContext
	translator
}

// Form represents an object that offers typical web form interaction.
type Form interface {
	Submit(context routeCtx) (isSubmitted bool, isValid bool, err error)
	GetUserID() string
	RememberMe() bool
	View() formview.View
	Close()
}

// Login web form.
type Login struct {
	validator validators.Validator
	userStore user.Store
	view      formview.View
	isClosed  bool
	userID    string
}

// Submit will submit and validate the form and handle all the rules.
func (login *Login) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
	if login.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = login.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = login.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = login.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = login.process(context); err != nil {
		return
	}
	return
}

func (login *Login) process(context translator) (bool, error) {
	user, ok, err := login.userStore.GetOneIDAndPasswordForActiveByEmail(login.validator.Field("email").Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !ok {
		login.validator.AddError(context.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err := login.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(login.validator.Field("password").Value())); err != nil {
		login.validator.AddError(context.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err = login.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	login.userID = user.ID()
	return true, nil
}

// GetUserID returns the User ID that was resolved while submitting the form.
func (login *Login) GetUserID() string {
	return login.userID
}

// RememberMe returns an indicator if the user wants to stay logged in longer.
func (login *Login) RememberMe() bool {
	return login.validator.Field("rememberMe").Value() == "1"
}

// View returns the FormView internal to help render inside html output.
func (login *Login) View() formview.View {
	return login.view
}

// Close will release internals.
func (login *Login) Close() {
	login.validator = nil
	login.userStore = nil
	login.view = nil
}
