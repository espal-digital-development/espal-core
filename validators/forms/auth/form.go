package auth

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

var _ Form = &Auth{}

type translator interface {
	Translate(string) string
}

type routeCtx interface {
	contexts.FormContext
	contexts.AuthenticationContext
	contexts.RequestContext
	translator
}

// Form represents an object that offers typical web form interaction.
type Form interface {
	Submit(routeCtx routeCtx) (isSubmitted bool, isValid bool, err error)
	GetUserID() string
	RememberMe() bool
	View() formview.View
	Close()
}

// Auth web form.
type Auth struct {
	validator validators.Validator
	userStore user.Store
	view      formview.View
	isClosed  bool
	userID    string
}

// Submit will submit and validate the form and handle all the rules.
func (a *Auth) Submit(routeCtx routeCtx) (isSubmitted bool, isValid bool, err error) {
	if a.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = a.validator.HandleFromRequest(routeCtx); err != nil {
		return
	}
	isSubmitted = a.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = a.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = a.process(routeCtx); err != nil {
		return
	}
	return
}

func (a *Auth) process(translator translator) (bool, error) {
	user, ok, err := a.userStore.GetOneActiveByEmail(a.validator.Field("email").Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !ok {
		a.validator.AddError(translator.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err := a.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	if isValid {
		hasAuthAccess, err := a.userStore.HasUserRight(user, "AccessAuth")
		if err != nil {
			return false, errors.Trace(err)
		}
		if !hasAuthAccess {
			a.validator.AddError(translator.Translate("theSuppliedCredentialsAreNotValid"))
		}
	}
	isValid, err = a.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password()),
		[]byte(a.validator.Field("password").Value())); err != nil {
		a.validator.AddError(translator.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err = a.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	a.userID = user.ID()
	return true, nil
}

// GetUserID returns the User ID that was resolved while submitting the form.
func (a *Auth) GetUserID() string {
	return a.userID
}

// RememberMe returns an indicator if the user wants to stay logged in longer.
func (a *Auth) RememberMe() bool {
	return a.validator.Field("rememberMe").Value() == "1"
}

// View returns the FormView internal to help render inside html output.
func (a *Auth) View() formview.View {
	return a.view
}

// Close will release internals.
func (a *Auth) Close() {
	a.validator = nil
	a.userStore = nil
	a.view = nil
}
