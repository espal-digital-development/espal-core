package validators

// ChoiceOption is used for select-fields options rendering.
type ChoiceOption interface {
	Value() string
	Display() string
}

type choiceOption struct {
	value   string
	display string
}

// Value returns the option value.
func (choiceOption *choiceOption) Value() string {
	return choiceOption.value
}

// Value returns the option display value.
func (choiceOption *choiceOption) Display() string {
	return choiceOption.display
}

// NewChoiceOption returns a new instance of ChoiceOption.
func (validators *Validators) NewChoiceOption(value string, display string) ChoiceOption {
	return &choiceOption{
		value:   value,
		display: display,
	}
}
