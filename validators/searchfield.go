package validators

// NewSearchField returns a new instance of FormField with the type Search.
func (f *Form) NewSearchField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: SearchFormField,
	})
}
