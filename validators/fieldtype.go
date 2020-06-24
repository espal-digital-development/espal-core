package validators

// FieldType determines the essence of a FormField.
type FieldType int8

// FormFields represents the Form-based field-types.
const (
	TokenFormField FieldType = iota + 1
	HoneypotFormField
	HiddenFormField
	TextFormField
	TextAreaFormField
	EmailFormField
	NumberFormField
	MoneyFormField
	PasswordFormField
	PercentageFormField
	SearchFormField
	URLFormField
	RangeFormField
	DateTimeFormField
	CheckboxFormField
	RadioFormField
	ChoiceFormField
	FileFormField
)
