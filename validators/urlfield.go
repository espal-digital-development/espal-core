package validators

// NewURLField returns a new instance of FormField with the type URL.
func (form *Form) NewURLField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: URLFormField,
	})
}
