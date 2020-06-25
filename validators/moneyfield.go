package validators

// NewMoneyField returns a new instance of FormField with the type Money.
func (f *Form) NewMoneyField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: MoneyFormField,
	})
}
