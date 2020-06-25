package validators

// NewHoneypotField returns a new instance of FormField with the type Honeypot.
func (f *Form) NewHoneypotField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: HoneypotFormField,
	})
}
