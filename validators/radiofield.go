package validators

// NewRadioField returns a new instance of FormField with the type Radio.
func (form *Form) NewRadioField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: RadioFormField,
	})
}
