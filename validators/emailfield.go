package validators

// NewEmailField returns a new instance of FormField with the type Email.
func (form *Form) NewEmailField(name string) FormField {
	return form.defaultChecks(&formField{
		name:      name,
		_type:     EmailFormField,
		minLength: 7,
		maxLength: 255,
	})
}
