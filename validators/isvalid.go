package validators

import (
	"fmt"
	"strings"

	"github.com/juju/errors"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

// IsValid checks and returns if all data rules are met.
// IsValid is not idempotent and might change when new errors are added to the validator.
// nolint:gocyclo,funlen
func (f *Form) IsValid() (bool, error) {
	isValid := false
	if len(f.submitErrors) > 0 {
		return isValid, nil
	}
	for _, field := range f.fields {
		if field.HasErrors() {
			return isValid, nil
		}
	}
	if f.isValidated {
		isValid = true
		return isValid, nil
	}
	f.isValidated = true

	var err error
	var tokenSupplied bool
	var honeypotSupplied bool

	// Trim all necessary fields first.
	// This needs to be done upfront because if field.NeedsToBeEqualToField needs to check
	// another field there might be a chance it wasn't trimmed yet and not validate correctly.
	for k := range f.fields {
		if f.fields[k].Trim() && f.fields[k].Value() != "" {
			f.fields[k].SetValue(strings.TrimSpace(f.fields[k].Value()))
		}
	}

	for _, field := range f.fields {
		// Reset the validation-errors
		field.RemoveAllErrors()

		if field.NeedsToBeEqualToField() != "" {
			targetField, ok := f.fields[field.NeedsToBeEqualToField()]
			if !ok {
				return false, errors.Errorf("target field `%s` for equal matching cannot be found", field.NeedsToBeEqualToField())
			}

			if field.Value() != targetField.Value() {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"theXFieldAndXFieldHaveToBeTheSame"), field.Name(), field.NeedsToBeEqualToField()))
			}
		}

		if field.CannotToBeEqualToField() != "" {
			targetField, ok := f.fields[field.CannotToBeEqualToField()]
			if !ok {
				return false, errors.Errorf("target field `%s` for -not equal- matching cannot be found",
					field.CannotToBeEqualToField())
			}

			if field.Value() == targetField.Value() {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"theXFieldAndXFieldCannotBeTheSame"), field.Name(), field.CannotToBeEqualToField()))
			}
		}

		switch field.Type() {
		case TokenFormField:
			tokenSupplied, err = f.validateTokenFormField(field)
			if err != nil {
				return false, errors.Trace(err)
			}
		case HoneypotFormField:
			honeypotSupplied = true
			if len(field.Value()) > 0 {
				return false, errors.Errorf("honeypot bait filled")
			}
		case PasswordFormField, HiddenFormField, TextFormField, TextAreaFormField, SearchFormField, EmailFormField:
			fieldLength := uint(len(field.Value()))
			if !field.Optional() && fieldLength == 0 {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXCannotBeEmpty"), field.Name()))
			}
			if field.Optional() && fieldLength == 0 {
				// Empty fields allowed on Optional and having a MinLength
			} else if field.MinLength() > 0 && fieldLength < field.MinLength() {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXHasToBeAtLeastXCharactersLong"), field.Name(), field.MinLength()))
			}
			if field.MaxLength() > 0 && fieldLength > field.MaxLength() {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXCannotBeLongerThanXCharacters"), field.Name(), field.MaxLength()))
			}
		case NumberFormField:
			err = f.validateNumberFormField(field)
			if err != nil {
				return false, errors.Trace(err)
			}
		case MoneyFormField:
			// TODO :: Implement
			return false, errors.Errorf("type MoneyFormField is not implemented yet")
		case PercentageFormField:
			// TODO :: Implement
			return false, errors.Errorf("type PercentageFormField is not implemented yet")
		case URLFormField:
			// TODO :: Implement
			return false, errors.Errorf("type URLFormField is not implemented yet")
		case RangeFormField:
			// TODO :: Implement
			return false, errors.Errorf("type RangeFormField is not implemented yet")
		case DateTimeFormField:
			err = f.validateDateFormField(field)
			if err != nil {
				return false, errors.Trace(err)
			}
		case CheckboxFormField:
			// No logic for now. Maybe never will have?
		case RadioFormField:
			// TODO :: Implement
			return false, errors.Errorf("type RadioFormField is not implemented yet")
		case ChoiceFormField:
			if err := f.validateChoiceFormField(field); err != nil {
				return false, errors.Trace(err)
			}
		case FileFormField:
			f.validateFileFormField(field)
		}

		// Validations
		switch field.Type() {
		case EmailFormField:
			if field.Optional() && len(field.Value()) == 0 {
				break
			}
			if field.Validate() && !f.regularExpressionsRepository.GetEmail().MatchString(field.Value()) {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXIsNotAValidEmail"), field.Name()))
			}
		case PasswordFormField:
			if field.Validate() && field.Value() != "" &&
				zxcvbn.PasswordStrength(field.Value(), make([]string, 0)).Score < 3 {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXPasswordIsNotSafeEnough"), field.Name()))
			}
		}
	}

	if !honeypotSupplied {
		f.submitErrors = append(f.submitErrors, f.translationsRepository.Singular(f.language.ID(),
			"honeypotFieldNotSupplied"))
	}
	if !tokenSupplied {
		f.submitErrors = append(f.submitErrors, f.translationsRepository.Singular(f.language.ID(),
			"validationTokenNotSupplied"))
	}

	if err := f.refreshToken(); err != nil {
		return false, errors.Trace(err)
	}
	if len(f.submitErrors) > 0 {
		return isValid, nil
	}

	for _, field := range f.fields {
		if field.HasErrors() {
			return isValid, nil
		}
	}

	isValid = true
	f.isValid = isValid
	return isValid, nil
}
