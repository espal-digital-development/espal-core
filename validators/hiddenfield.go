package validators

// NewHiddenField returns a new instance of FormField with the type Hidden.
func (f *Form) NewHiddenField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: HiddenFormField,
	})
}
