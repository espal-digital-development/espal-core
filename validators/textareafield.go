package validators

// NewTextAreaField returns a new instance of FormField with the type TextArea.
func (form *Form) NewTextAreaField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: TextAreaFormField,
	})
}
