package validators

// NewChoiceField returns a new instance of FormField with the type Choice.
func (form *Form) NewChoiceField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: ChoiceFormField,
	})
}

// SetSearchable marks the form field that options are searchable.
func (formField *formField) SetSearchable() {
	formField.searchable = true
}

// Searchable returns if the field is searchable.
func (formField *formField) Searchable() bool {
	return formField.searchable
}

// SetSearchableDataPath sets the form field data path the search should fetch from.
func (formField *formField) SetSearchableDataPath(path string) {
	formField.searchableDataPath = path
}

// SearchableDataPath returns the form field data path that will be searched from.
func (formField *formField) SearchableDataPath() string {
	return formField.searchableDataPath
}

// DetermineNoSelectText sets the no-selection text based on it's value.
func (formField *formField) DetermineNoSelectText(localeID uint16) {
	// TODO :: 777 Fix to make this actually facilitate translations somehow
	// if formField.value == "" {
	// 	formField.NoSelectionText = formField.translations.Singular(localeID, "makeAChoice")
	// } else {
	// 	formField.NoSelectionText = formField.translations.Singluar(localeID, "clearSelection")
	// }
}

// SetCheckValuesInChoices will check if the choices that are submitted are actually
// present in the options list of the form, not allowing dynamic new values.
func (formField *formField) SetCheckValuesInChoices() {
	formField.checkValuesInChoices = true
}

// CheckValuesInChoices returns if choices should be checked.
func (formField *formField) CheckValuesInChoices() bool {
	return formField.checkValuesInChoices
}

// SetMultiple makes the form field able to receive multiple values.
func (formField *formField) SetMultiple() {
	formField.multiple = true
}

// Multiple returns if the formfield can have multiple choice options.
func (formField *formField) Multiple() bool {
	return formField.multiple
}

// SetChoices sets the form field choices.
func (formField *formField) SetChoices(choices []ChoiceOption) {
	formField.choices = choices
}

// AddChoice adds to the form field choices.
func (formField *formField) AddChoice(choice ChoiceOption) {
	formField.choices = append(formField.choices, choice)
}

// ChoiceIsSelected checks if the given option is currently selected.
func (formField *formField) ChoiceIsSelected(choiceOption ChoiceOption) bool {
	if formField.multiple {
		for k := range formField.values {
			if formField.values[k] == choiceOption.Value() {
				return true
			}
		}
	} else if choiceOption.Value() == formField.Value() {
		return true
	}

	return false
}

// Choices returns the choices set on the field.
func (formField *formField) Choices() []ChoiceOption {
	return formField.choices
}
