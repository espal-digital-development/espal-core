package validators

// NewTokenField returns a new instance of FormField with the type Token.
func (form *Form) NewTokenField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: TokenFormField,
	})
}
