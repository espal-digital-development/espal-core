package validators

// NewRadioField returns a new instance of FormField with the type Radio.
func (f *Form) NewRadioField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: RadioFormField,
	})
}
