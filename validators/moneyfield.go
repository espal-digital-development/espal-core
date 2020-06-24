package validators

// NewMoneyField returns a new instance of FormField with the type Money.
func (form *Form) NewMoneyField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: MoneyFormField,
	})
}
