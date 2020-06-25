package validators

// NewURLField returns a new instance of FormField with the type URL.
func (f *Form) NewURLField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: URLFormField,
	})
}
