package validators

import (
	"fmt"
	"strings"

	"github.com/juju/errors"
)

// nolint:gocyclo
func (f *Form) validateChoiceFormField(field *formField) error {
	var values []string
	fieldValuesLength := len(field.Values())
	if !field.Multiple() && fieldValuesLength > 0 {
		field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
			"fieldXDoesNotAllowMultipleValues"), field.Name()))
	}

	if field.Multiple() && fieldValuesLength > 0 && field.Value() == "" {
		values = field.Values()
	} else if field.Value() != "" {
		values = append(values, field.Value())
	}

	valuesLength := len(values)
	if !field.Optional() && valuesLength == 0 {
		field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
			"fieldXCannotBeEmpty"), field.Name()))
	}
	if field.Multiple() && valuesLength == 1 && field.Value() != "" {
		// Favor filling field.values over multiple with one value filling field.value
		field.SetValue("")
		if err := field.SetValues(values); err != nil {
			return errors.Trace(err)
		}
	}
	if len(field.Values()) > 0 {
		if field.Trim() {
			values := field.Values()
			for k := range values {
				values[k] = strings.TrimSpace(values[k])
			}
			if err := field.SetValues(values); err != nil {
				return errors.Trace(err)
			}
		}

		allowedValuesCount := len(field.AllowedValues())
		disallowedValuesCount := len(field.DisallowValues())
		if allowedValuesCount > 0 || disallowedValuesCount > 0 {
			for _, value := range field.Values() {
				var allowed bool
				if allowedValuesCount > 0 {
					for _, allowedValue := range field.AllowedValues() {
						if allowedValue == value {
							allowed = true
							break
						}
					}
				} else {
					allowed = true
				}
				var disallowed bool
				for _, disallowedValue := range field.DisallowValues() {
					if disallowedValue == value {
						disallowed = true
						break
					}
				}
				if !allowed || disallowed {
					field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
						"valueXIsNotAllowedForFieldX"), value, field.Name()))
				}
			}
		}
	}

	if field.CheckValuesInChoices() && (field.Value() != "" || len(field.Values()) > 0) {
		if len(field.Values()) > 0 {
			var found int
			for _, choice := range field.Choices() {
				for _, value := range field.Values() {
					if choice.Value() == value {
						found++
					}
				}
			}
			if found < len(field.Values()) {
				field.AddError(fmt.Sprintf(f.translationsRepository.FormattedPlural(f.language.ID(),
					"fieldXContainsInvalidChosenValue"), field.Name()))
			}
		} else {
			var valid bool
			for _, choice := range field.Choices() {
				if choice.Value() == field.Value() {
					valid = true
					break
				}
			}
			if !valid {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXContainsInvalidChosenValue"), field.Name()))
			}
		}
	}

	field.DetermineNoSelectText(f.language.ID())

	return nil
}
