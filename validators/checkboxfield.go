package validators

// NewCheckboxField returns a new instance of FormField with the type Checkbox.
func (form *Form) NewCheckboxField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: CheckboxFormField,
	})
}
