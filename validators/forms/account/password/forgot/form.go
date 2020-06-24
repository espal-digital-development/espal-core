package forgot

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Form = &Forgot{}

type context interface {
	contexts.RequestContext
	contexts.FormContext
}

// Form represents an object that offers typical web form interaction.
type Form interface {
	Submit(context context) (isSubmitted bool, isValid bool, err error)
	FormFieldValue(name string) string
	View() formview.View
	Close()
}

// Forgot web form.
type Forgot struct {
	validator validators.Validator
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (forgot *Forgot) Submit(context context) (isSubmitted bool, isValid bool, err error) {
	if forgot.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = forgot.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = forgot.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = forgot.validator.IsValid()
		if err != nil {
			return
		}
	}
	return
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (forgot *Forgot) FormFieldValue(name string) string {
	return forgot.validator.FieldValue(name)
}

// View returns the FormView internal to help render inside html output.
func (forgot *Forgot) View() formview.View {
	return forgot.view
}

// Close will release internals.
func (forgot *Forgot) Close() {
	forgot.validator = nil
	forgot.view = nil
}
