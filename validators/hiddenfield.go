package validators

// NewHiddenField returns a new instance of FormField with the type Hidden.
func (form *Form) NewHiddenField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: HiddenFormField,
	})
}
