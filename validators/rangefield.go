package validators

// NewRangeField returns a new instance of FormField with the type Range.
func (f *Form) NewRangeField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: RangeFormField,
	})
}
