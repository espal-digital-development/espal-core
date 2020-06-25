package validators

// NewTextField returns a new instance of FormField with the type Text.
func (f *Form) NewTextField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: TextFormField,
	})
}
