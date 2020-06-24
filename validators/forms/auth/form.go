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
func (auth *Auth) Submit(routeCtx routeCtx) (isSubmitted bool, isValid bool, err error) {
	if auth.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = auth.validator.HandleFromRequest(routeCtx); err != nil {
		return
	}
	isSubmitted = auth.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = auth.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = auth.process(routeCtx); err != nil {
		return
	}
	return
}

func (auth *Auth) process(translator translator) (bool, error) {
	user, ok, err := auth.userStore.GetOneActiveByEmail(auth.validator.Field("email").Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !ok {
		auth.validator.AddError(translator.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err := auth.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	if isValid {
		hasAuthAccess, err := auth.userStore.HasUserRight(user, "AccessAuth")
		if err != nil {
			return false, errors.Trace(err)
		}
		if !hasAuthAccess {
			auth.validator.AddError(translator.Translate("theSuppliedCredentialsAreNotValid"))
		}
	}
	isValid, err = auth.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(auth.validator.Field("password").Value())); err != nil {
		auth.validator.AddError(translator.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err = auth.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	auth.userID = user.ID()
	return true, nil
}

// GetUserID returns the User ID that was resolved while submitting the form.
func (auth *Auth) GetUserID() string {
	return auth.userID
}

// RememberMe returns an indicator if the user wants to stay logged in longer.
func (auth *Auth) RememberMe() bool {
	return auth.validator.Field("rememberMe").Value() == "1"
}

// View returns the FormView internal to help render inside html output.
func (auth *Auth) View() formview.View {
	return auth.view
}

// Close will release internals.
func (auth *Auth) Close() {
	auth.validator = nil
	auth.userStore = nil
	auth.view = nil
}
