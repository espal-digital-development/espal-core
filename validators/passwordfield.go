package validators

// NewPasswordField returns a new instance of FormField with the type Password.
func (form *Form) NewPasswordField(name string) FormField {
	return form.defaultChecks(&formField{
		name:      name,
		_type:     PasswordFormField,
		minLength: 5,
		maxLength: 72,
	})
}
