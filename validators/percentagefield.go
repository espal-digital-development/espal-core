package validators

// NewPercentageField returns a new instance of FormField with the type Percentage.
func (f *Form) NewPercentageField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: PercentageFormField,
	})
}
