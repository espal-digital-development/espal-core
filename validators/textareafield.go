package validators

// NewTextAreaField returns a new instance of FormField with the type TextArea.
func (f *Form) NewTextAreaField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: TextAreaFormField,
	})
}
