package validators

// NewPercentageField returns a new instance of FormField with the type Percentage.
func (form *Form) NewPercentageField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: PercentageFormField,
	})
}
