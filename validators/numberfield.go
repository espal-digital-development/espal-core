package validators

// NewNumberField returns a new instance of FormField with the type Number.
func (form *Form) NewNumberField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: NumberFormField,
	})
}

// SetMinValue sets the minimal numeral value.
func (formField *formField) SetMinValue(minValue float64) {
	formField.minValue = minValue
}

// MinValue gets the minimal numeral value.
func (formField *formField) MinValue() float64 {
	return formField.minValue
}

// SetMaxValue sets the maximum numeral value.
func (formField *formField) SetMaxValue(maxValue float64) {
	formField.maxValue = maxValue
}

// MaxValue gets the maximum numeral value.
func (formField *formField) MaxValue() float64 {
	return formField.maxValue
}
