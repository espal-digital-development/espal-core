package validators

import (
	"strconv"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/storage"
	"github.com/espal-digital-development/espal-core/tokenpool"
	"github.com/juju/errors"
)

var _ Validator = &Form{}

// TODO :: 7777 Maybe can split the Validator interface methods up into smaller ones for more flexibility.

// Validator represents an object that can validate data before
// it gets processed.
type Validator interface {
	LoadViewData() Validator
	SetPostloadCallback(callback func(Validator))
	SetPreSubmitCallback(callback func(Validator))
	SetPostSubmitCallback(callback func(Validator))
	SetPreViewCallback(callback func(Validator))
	ContainsSelectSearch() bool
	AddError(errorString string)
	Field(name string) FormField
	FieldValue(name string) string
	FieldExists(name string) (exists bool)
	ClearFields() error

	HandleFromRequest(context context) error
	IsSubmitted() bool
	IsValid() (bool, error)

	AddField(field FormField) error

	NewTokenField(name string) FormField
	NewHoneypotField(name string) FormField
	NewHiddenField(name string) FormField
	NewTextField(name string) FormField
	NewTextAreaField(name string) FormField
	NewEmailField(name string) FormField
	NewNumberField(name string) FormField
	NewMoneyField(name string) FormField
	NewPasswordField(name string) FormField
	NewPercentageField(name string) FormField
	NewSearchField(name string) FormField
	NewURLField(name string) FormField
	NewRangeField(name string) FormField
	NewDateTimeField(name string) FormField
	NewCheckboxField(name string) FormField
	NewRadioField(name string) FormField
	NewChoiceField(name string) FormField
	NewFileField(name string, storage storage.Storage) FormField

	RenderErrors() string
	RenderOpen() string
	RenderField(name string) string
	RenderCreateUpdateActions(fieldName string, url string) string
}

// Form validator for web-based forms.
type Form struct {
	configService                config.Config
	loggerService                logger.Loggable
	translationsRepository       translations.Repository
	regularExpressionsRepository regularexpressions.Repository
	tokenPoolService             tokenpool.Pool

	isValid         bool
	isValidated     bool
	isSubmitted     bool
	isFormValidator bool
	submitErrors    []string
	fields          map[string]*formField
	language        language

	// Callback Flow:
	// NewForm() is called
	// postLoad is called
	// preSubmit is called if the method is POST
	// HandleFromRequest() is triggered (submit)
	// postSubmit is called if the method is POST
	// preView is called
	postloadCallback   func(Validator)
	preSubmitCallback  func(Validator)
	postSubmitCallback func(Validator)
	preViewCallback    func(Validator)
}

// LoadViewData will explicitly trigger the preViewCallback and return the instance.
func (form *Form) LoadViewData() Validator {
	if form.preViewCallback != nil {
		form.preViewCallback(form)
	}
	return form
}

// SetPostloadCallback will call the callback function before the Form is handled in initial request.
// In a typical Form request-cycle this will be the initial (GET) load of the Form.
func (form *Form) SetPostloadCallback(callback func(Validator)) {
	form.postloadCallback = callback
}

// SetPreSubmitCallback will call the callback function before the Form submit is processed in the posted request.
// In a typical Form request-cycle this will be the pre-submit (POST) of the Form.
func (form *Form) SetPreSubmitCallback(callback func(Validator)) {
	form.preSubmitCallback = callback
}

// SetPostSubmitCallback will call the callback function after the Form submit is processed in the posted request.
// In a typical Form request-cycle this will be the post-submit (POST) of the Form.
func (form *Form) SetPostSubmitCallback(callback func(Validator)) {
	form.postSubmitCallback = callback
}

// SetPreViewCallback will call the callback function when the .LoadViewData() function is called.
func (form *Form) SetPreViewCallback(callback func(Validator)) {
	form.preViewCallback = callback
}

// ContainsSelectSearch determines if at least one ChoiceType field is
// present with Searchable active.
func (form *Form) ContainsSelectSearch() bool {
	for _, field := range form.fields {
		if ChoiceFormField == field.Type() && field.Searchable() {
			return true
		}
	}
	return false
}

// IsSubmitted checks and returns if the form was submitted.
func (form *Form) IsSubmitted() bool {
	return form.isSubmitted
}

// AddError adds a submit-error to the submitErrors stack.
func (form *Form) AddError(errorString string) {
	form.submitErrors = append(form.submitErrors, errorString)
	form.isValid = false
}

func (form *Form) field(name string) *formField {
	if !form.isFormValidator {
		err := errors.Errorf("cannot call Field on a non-form validator")
		form.loggerService.Error(err.Error())
		panic(err)
	}
	field, ok := form.fields[name]
	if !ok {
		err := errors.Errorf("field %s couldn't be found", name)
		form.loggerService.Error(err.Error())
		panic(err)
	}
	return field
}

// Field returns the definition of a given field from the Form object.
// Keeping the panics here because it's chained constantly.
func (form *Form) Field(name string) FormField {
	return form.field(name)
}

// FieldValue instantly returns the value of the given field.
func (form *Form) FieldValue(name string) string {
	return form.Field(name).Value()
}

// FieldExists returns if the requested field exists
func (form *Form) FieldExists(name string) (exists bool) {
	if _, ok := form.fields[name]; ok {
		exists = true
	}
	return
}

// AddField adds a new field to the form.
func (form *Form) AddField(field FormField) error {
	if field.Name() == "" {
		return errors.Errorf("field should have a name")
	}
	if form.FieldExists(field.Name()) {
		return errors.Errorf("field %s already exists", field.Name())
	}
	form.fields[field.Name()] = field.(*formField)
	return nil
}

// ClearFields clears all fields except the token and honeypot.
// The validity of the form will automatically be disqualified too.
func (form *Form) ClearFields() error {
	form.isValid = false
	form.isValidated = false
	for _, field := range form.fields {
		if TokenFormField == field.Type() || HoneypotFormField == field.Type() {
			continue
		}
		if ChoiceFormField == field.Type() {
			if err := field.SetValues(make([]string, 0)); err != nil {
				return errors.Trace(err)
			}
		}
		field.SetValue("")
	}
	return nil
}

func (form *Form) isMultipart() bool {
	for _, field := range form.fields {
		if FileFormField == field.Type() {
			return true
		}
	}
	return false
}

func (form *Form) refreshToken() error {
	field := form.Field("_t")
	oldToken, err := strconv.Atoi(field.Value())
	if err != nil {
		return errors.Trace(err)
	}
	form.tokenPoolService.Expire(oldToken)
	token, err := form.tokenPoolService.RequestToken()
	if err != nil {
		return errors.Trace(err)
	}
	field.SetValue(strconv.Itoa(token))
	return nil
}

func (form *Form) defaultChecks(formField *formField) *formField {
	if formField.Name() == "" {
		formField.AddError("nameless fields aren't allowed")
		return formField
	}

	return formField
}

func (form *Form) perror(i int, err error) {
	if err != nil {
		form.loggerService.Error(errors.ErrorStack(err))
	}
}
