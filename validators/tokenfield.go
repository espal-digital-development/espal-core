package validators

// NewTokenField returns a new instance of FormField with the type Token.
func (f *Form) NewTokenField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: TokenFormField,
	})
}
