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
func (o *choiceOption) Value() string {
	return o.value
}

// Value returns the option display value.
func (o *choiceOption) Display() string {
	return o.display
}

// NewChoiceOption returns a new instance of ChoiceOption.
func (v *Validators) NewChoiceOption(value string, display string) ChoiceOption {
	return &choiceOption{
		value:   value,
		display: display,
	}
}
