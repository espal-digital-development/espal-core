package validators

const (
	emailFieldMinLength = 7
	emailFieldMaxLength = 255
)

// NewEmailField returns a new instance of FormField with the type Email.
func (f *Form) NewEmailField(name string) FormField {
	return f.defaultChecks(&formField{
		name:      name,
		_type:     EmailFormField,
		minLength: emailFieldMinLength,
		maxLength: emailFieldMaxLength,
	})
}
