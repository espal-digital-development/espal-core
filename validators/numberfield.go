package validators

// NewNumberField returns a new instance of FormField with the type Number.
func (f *Form) NewNumberField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: NumberFormField,
	})
}

// SetMinValue sets the minimal numeral value.
func (f *formField) SetMinValue(minValue float64) {
	f.minValue = minValue
}

// MinValue gets the minimal numeral value.
func (f *formField) MinValue() float64 {
	return f.minValue
}

// SetMaxValue sets the maximum numeral value.
func (f *formField) SetMaxValue(maxValue float64) {
	f.maxValue = maxValue
}

// MaxValue gets the maximum numeral value.
func (f *formField) MaxValue() float64 {
	return f.maxValue
}
