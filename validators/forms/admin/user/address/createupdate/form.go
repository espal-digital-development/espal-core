package createupdate

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Form = &CreateUpdate{}

type context interface {
	contexts.RequestContext
	contexts.FormContext
}

// Form represents an object that offers typical web form interaction.
type Form interface {
	Submit(context context) (isSubmitted bool, isValid bool, err error)
	View() formview.View
	FieldValue(name string) string
	FieldPointerValue(name string) *string
	FieldValues(name string) []string
	FieldValueAsBool(name string) bool
	FieldValueAsUint16(name string) uint16
	Close()
}

// CreateUpdate web form.
type CreateUpdate struct {
	validator validators.Validator
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (createUpdate *CreateUpdate) Submit(context context) (isSubmitted bool, isValid bool, err error) {
	if createUpdate.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = createUpdate.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = createUpdate.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = createUpdate.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	return
}

// View returns the FormView internal to help render inside html output.
func (createUpdate *CreateUpdate) View() formview.View {
	return createUpdate.view
}

// FieldValue returns a single form value that was resolved while submitting the form.
func (createUpdate *CreateUpdate) FieldValue(name string) string {
	return createUpdate.validator.FieldValue(name)
}

// FieldPointerValue returns a single form pointer value that was resolved while submitting the form.
func (createUpdate *CreateUpdate) FieldPointerValue(name string) *string {
	return createUpdate.validator.Field(name).PointerValue()
}

// FieldValues returns a single form values that was resolved while submitting the form.
func (createUpdate *CreateUpdate) FieldValues(name string) []string {
	return createUpdate.validator.Field(name).Values()
}

// FieldValueAsBool returns a single form bool value that was resolved while submitting the form.
func (createUpdate *CreateUpdate) FieldValueAsBool(name string) bool {
	return createUpdate.validator.Field(name).ValueAsBool()
}

// FieldValueAsUint16 returns a single form uint16 value that was resolved while submitting the form.
func (createUpdate *CreateUpdate) FieldValueAsUint16(name string) uint16 {
	return createUpdate.validator.Field(name).ValueAsUint16()
}

// Close will release internals.
func (createUpdate *CreateUpdate) Close() {
	createUpdate.validator = nil
	createUpdate.view = nil
}
