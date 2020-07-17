package validators

// NewChoiceField returns a new instance of FormField with the type Choice.
func (f *Form) NewChoiceField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: ChoiceFormField,
	})
}

// SetSearchable marks the form field that options are searchable.
func (f *formField) SetSearchable() {
	f.searchable = true
}

// Searchable returns if the field is searchable.
func (f *formField) Searchable() bool {
	return f.searchable
}

// SetSearchableDataPath sets the form field data path the search should fetch from.
func (f *formField) SetSearchableDataPath(path string) {
	f.searchableDataPath = path
}

// SearchableDataPath returns the form field data path that will be searched from.
func (f *formField) SearchableDataPath() string {
	return f.searchableDataPath
}

// DetermineNoSelectText sets the no-selection text based on it's value.
func (f *formField) DetermineNoSelectText(localeID uint16) {
	// TODO :: 777 Fix to make this actually facilitate translations somehow
	// if  f.value == "" {
	// 	 f.NoSelectionText =  f.translations.Singular(localeID, "makeAChoice")
	// } else {
	// 	 f.NoSelectionText =  f.translations.Singluar(localeID, "clearSelection")
	// }
}

// SetCheckValuesInChoices will check if the choices that are submitted are actually present in the options list of the
// form, not allowing dynamic new values.
func (f *formField) SetCheckValuesInChoices() {
	f.checkValuesInChoices = true
}

// CheckValuesInChoices returns if choices should be checked.
func (f *formField) CheckValuesInChoices() bool {
	return f.checkValuesInChoices
}

// SetMultiple makes the form field able to receive multiple values.
func (f *formField) SetMultiple() {
	f.multiple = true
}

// Multiple returns if the formfield can have multiple choice options.
func (f *formField) Multiple() bool {
	return f.multiple
}

// SetChoices sets the form field choices.
func (f *formField) SetChoices(choices []ChoiceOption) {
	f.choices = choices
}

// AddChoice adds to the form field choices.
func (f *formField) AddChoice(choice ChoiceOption) {
	f.choices = append(f.choices, choice)
}

// ChoiceIsSelected checks if the given option is currently selected.
func (f *formField) ChoiceIsSelected(choiceOption ChoiceOption) bool {
	if f.multiple {
		for k := range f.values {
			if f.values[k] == choiceOption.Value() {
				return true
			}
		}
	} else if choiceOption.Value() == f.Value() {
		return true
	}

	return false
}

// Choices returns the choices set on the field.
func (f *formField) Choices() []ChoiceOption {
	return f.choices
}
