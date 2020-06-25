package validators

// NewPasswordField returns a new instance of FormField with the type Password.
func (f *Form) NewPasswordField(name string) FormField {
	return f.defaultChecks(&formField{
		name:      name,
		_type:     PasswordFormField,
		minLength: 5,
		maxLength: 72,
	})
}
