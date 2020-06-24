package validators

// NewSearchField returns a new instance of FormField with the type Search.
func (form *Form) NewSearchField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: SearchFormField,
	})
}
