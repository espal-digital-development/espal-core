package validators

// NewCheckboxField returns a new instance of FormField with the type Checkbox.
func (f *Form) NewCheckboxField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: CheckboxFormField,
	})
}
