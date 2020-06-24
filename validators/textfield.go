package validators

// NewTextField returns a new instance of FormField with the type Text.
func (form *Form) NewTextField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: TextFormField,
	})
}
