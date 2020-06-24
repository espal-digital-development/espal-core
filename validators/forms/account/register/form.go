package register

import (
	"database/sql"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Form = &Register{}

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
	FormFieldValue(name string) string
	View() formview.View
	Close()
}

// Register web form.
type Register struct {
	validator validators.Validator
	userStore user.Store
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (register *Register) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
	if register.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = register.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = register.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = register.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = register.process(context); err != nil {
		return
	}
	return
}

func (register *Register) process(translator translator) (bool, error) {
	exists, err := register.userStore.ExistsByEmail(register.validator.Field("email").Value())
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Trace(err)
	}
	if exists {
		register.validator.Field("email").AddError(translator.Translate("emailIsAlreadyUsed"))
	}
	return register.validator.IsValid()
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (register *Register) FormFieldValue(name string) string {
	return register.validator.FieldValue(name)
}

// View returns the FormView internal to help render inside html output.
func (register *Register) View() formview.View {
	return register.view
}

// Close will release internals.
func (register *Register) Close() {
	register.validator = nil
	register.userStore = nil
	register.view = nil
}
