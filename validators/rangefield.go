package validators

// NewRangeField returns a new instance of FormField with the type Range.
func (form *Form) NewRangeField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: RangeFormField,
	})
}
