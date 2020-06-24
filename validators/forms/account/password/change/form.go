package change

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

var _ Form = &Change{}

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
	View() formview.View
	FormFieldValue(name string) string
	Close()
}

// Change web form.
type Change struct {
	validator validators.Validator
	user      user.UserEntity
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (change *Change) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
	if change.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = change.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = change.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = change.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = change.process(context); err != nil {
		return
	}
	return
}

func (change *Change) process(translator translator) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(change.user.Password()), []byte(change.validator.Field("currentPassword").Value())); err != nil {
		change.validator.AddError(translator.Translate("yourCurrentPasswordDidNotMatch"))
	}
	isValid, err := change.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	return true, nil
}

// View returns the FormView internal to help render inside html output.
func (change *Change) View() formview.View {
	return change.view
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (change *Change) FormFieldValue(name string) string {
	return change.validator.FieldValue(name)
}

// Close will release internals.
func (change *Change) Close() {
	change.validator = nil
	change.view = nil
}
