package recovery

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

var _ Form = &Recovery{}

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
	GetPasswordResetCount() *uint8
	FormFieldValue(name string) string
	View() formview.View
	Close()
}

// Recovery web form.
type Recovery struct {
	validator          validators.Validator
	userStore          user.Store
	view               formview.View
	isClosed           bool
	userID             string
	passwordResetCount *uint8
}

// Submit will submit and validate the form and handle all the rules.
func (recovery *Recovery) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
	if recovery.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = recovery.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = recovery.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = recovery.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = recovery.process(context); err != nil {
		return
	}
	return
}

func (recovery *Recovery) process(translator translator) (bool, error) {
	user, ok, err := recovery.userStore.GetOneByEmail(recovery.validator.Field("email").Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !ok {
		recovery.validator.AddError(translator.Translate("theSuppliedInformationDoesNotMatchAnyActiveAccount"))
		return false, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(recovery.validator.Field("newPassword").Value())); err == nil {
		recovery.validator.AddError(translator.Translate("yourPasswordShouldNotBeTheSameAsBefore"))
		return false, nil
	}
	recovery.userID = user.ID()
	recovery.passwordResetCount = user.PasswordResetCount()
	return true, nil
}

// GetUserID returns the User ID that was resolved while submitting the form.
func (recovery *Recovery) GetUserID() string {
	return recovery.userID
}

// GetPasswordResetCount returns the User PasswordResetCount that was resolved while submitting the form.
func (recovery *Recovery) GetPasswordResetCount() *uint8 {
	return recovery.passwordResetCount
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (recovery *Recovery) FormFieldValue(name string) string {
	return recovery.validator.FieldValue(name)
}

// View returns the FormView internal to help render inside html output.
func (recovery *Recovery) View() formview.View {
	return recovery.view
}

// Close will release internals.
func (recovery *Recovery) Close() {
	recovery.validator = nil
	recovery.userStore = nil
	recovery.view = nil
}
