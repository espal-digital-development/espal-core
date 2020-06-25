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
func (f *Form) LoadViewData() Validator {
	if f.preViewCallback != nil {
		f.preViewCallback(f)
	}
	return f
}

// SetPostloadCallback will call the callback function before the Form is handled in initial request.
// In a typical Form request-cycle this will be the initial (GET) load of the Form.
func (f *Form) SetPostloadCallback(callback func(Validator)) {
	f.postloadCallback = callback
}

// SetPreSubmitCallback will call the callback function before the Form submit is processed in the posted request.
// In a typical Form request-cycle this will be the pre-submit (POST) of the Form.
func (f *Form) SetPreSubmitCallback(callback func(Validator)) {
	f.preSubmitCallback = callback
}

// SetPostSubmitCallback will call the callback function after the Form submit is processed in the posted request.
// In a typical Form request-cycle this will be the post-submit (POST) of the Form.
func (f *Form) SetPostSubmitCallback(callback func(Validator)) {
	f.postSubmitCallback = callback
}

// SetPreViewCallback will call the callback function when the .LoadViewData() function is called.
func (f *Form) SetPreViewCallback(callback func(Validator)) {
	f.preViewCallback = callback
}

// ContainsSelectSearch determines if at least one ChoiceType field is
// present with Searchable active.
func (f *Form) ContainsSelectSearch() bool {
	for _, field := range f.fields {
		if ChoiceFormField == field.Type() && field.Searchable() {
			return true
		}
	}
	return false
}

// IsSubmitted checks and returns if the form was submitted.
func (f *Form) IsSubmitted() bool {
	return f.isSubmitted
}

// AddError adds a submit-error to the submitErrors stack.
func (f *Form) AddError(errorString string) {
	f.submitErrors = append(f.submitErrors, errorString)
	f.isValid = false
}

func (f *Form) field(name string) *formField {
	if !f.isFormValidator {
		err := errors.Errorf("cannot call Field on a non-form validator")
		f.loggerService.Error(err.Error())
		panic(err)
	}
	field, ok := f.fields[name]
	if !ok {
		err := errors.Errorf("field %s couldn't be found", name)
		f.loggerService.Error(err.Error())
		panic(err)
	}
	return field
}

// Field returns the definition of a given field from the Form object.
// Keeping the panics here because it's chained constantly.
func (f *Form) Field(name string) FormField {
	return f.field(name)
}

// FieldValue instantly returns the value of the given field.
func (f *Form) FieldValue(name string) string {
	return f.Field(name).Value()
}

// FieldExists returns if the requested field exists
func (f *Form) FieldExists(name string) (exists bool) {
	if _, ok := f.fields[name]; ok {
		exists = true
	}
	return
}

// AddField adds a new field to the form.
func (f *Form) AddField(field FormField) error {
	if field.Name() == "" {
		return errors.Errorf("field should have a name")
	}
	if f.FieldExists(field.Name()) {
		return errors.Errorf("field %s already exists", field.Name())
	}
	f.fields[field.Name()] = field.(*formField)
	return nil
}

// ClearFields clears all fields except the token and honeypot.
// The validity of the form will automatically be disqualified too.
func (f *Form) ClearFields() error {
	f.isValid = false
	f.isValidated = false
	for _, field := range f.fields {
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

func (f *Form) isMultipart() bool {
	for _, field := range f.fields {
		if FileFormField == field.Type() {
			return true
		}
	}
	return false
}

func (f *Form) refreshToken() error {
	field := f.Field("_t")
	oldToken, err := strconv.Atoi(field.Value())
	if err != nil {
		return errors.Trace(err)
	}
	f.tokenPoolService.Expire(oldToken)
	token, err := f.tokenPoolService.RequestToken()
	if err != nil {
		return errors.Trace(err)
	}
	field.SetValue(strconv.Itoa(token))
	return nil
}

func (f *Form) defaultChecks(formField *formField) *formField {
	if formField.Name() == "" {
		f.AddError("nameless fields aren't allowed")
		return formField
	}

	return formField
}

func (f *Form) perror(i int, err error) {
	if err != nil {
		f.loggerService.Error(errors.ErrorStack(err))
	}
}
