package validators

// NewHoneypotField returns a new instance of FormField with the type Honeypot.
func (form *Form) NewHoneypotField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: HoneypotFormField,
	})
}
