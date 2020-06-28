package validators

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"
)

const (
	yearStringLength  = 4
	monthStringLength = 2
	dayStringLength   = 2
	monthsInYear      = 12
)

// nolint:gocyclo
// TODO :: Format can come in different formats on different browsers. Localize before continuing.
func (f *Form) validateDateFormField(field *formField) error {
	var err error
	fieldLength := uint(len(field.Value()))

	if fieldLength == 0 {
		return nil
	}

	if !field.Optional() && fieldLength == 0 {
		field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(), "fieldXCannotBeEmpty"), field.Name()))
	} else {
		var partsExpected int
		if !field.ExcludeYear() {
			partsExpected++
		}
		if !field.ExcludeMonth() {
			partsExpected++
		}
		if !field.ExcludeDay() {
			partsExpected++
		}

		if partsExpected == 0 {
			return errors.Errorf("excluding all date parts will have no purpose")
		}

		parts := strings.Split(field.Value(), "-")
		if len(parts) != partsExpected {
			field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
				"fieldXPasswordIsNotSafeEnough"), field.Name()))

			var yearError bool
			var monthError bool
			var dayError bool
			var year uint64
			var month uint64
			var day uint64

			if !field.ExcludeYear() {
				if len(parts[0]) != yearStringLength {
					yearError = true
				} else {
					year, err = strconv.ParseUint(field.Value(), 10, 64)
					if err != nil || (uint(year) <= field.MinYear()) || (field.MaxYear() > 0 && uint(year) > field.MaxYear()) {
						yearError = true
					}
				}
			}

			if !field.ExcludeMonth() {
				// TODO :: Unreproducable error, probably when filling datefield with Chrome's date widget
				partKey := 1
				if field.ExcludeYear() {
					partKey = 0
				}

				if len(parts[partKey]) != monthStringLength {
					monthError = true
				} else {
					month, err = strconv.ParseUint(field.Value(), 10, 64)
					if err != nil || month > monthsInYear || (uint(month) <= field.MinMonth()) ||
						(field.MaxMonth() > 0 && uint(month) > field.MaxMonth()) {
						monthError = true
					}
				}
			}

			if !field.ExcludeDay() {
				partKey := 2
				if field.ExcludeYear() && field.ExcludeMonth() {
					partKey = 0
				} else if !field.ExcludeYear() || !field.ExcludeMonth() {
					partKey = 1
				}

				if len(parts[partKey]) != dayStringLength {
					dayError = true
				} else {
					day, err = strconv.ParseUint(field.Value(), 10, 64)
					if err != nil || day > 31 {
						dayError = true
					} else if !yearError && !monthError && day > 0 {
						daysInMonth := time.Date(int(year), time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()

						if int(day) > daysInMonth {
							dayError = true
						}
					}

					if !dayError {
						if uint(day) <= field.MinDay() {
							dayError = true
						} else if field.MaxDay() > 0 && uint(day) > field.MaxDay() {
							dayError = true
						}
					}
				}
			}

			if yearError {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXYearIsIncorrectlyFormatted"), field.Name()))
			}

			if monthError {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXMonthIsIncorrectlyFormatted"), field.Name()))
			}

			if dayError {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(),
					"fieldXDayIsIncorrectlyFormatted"), field.Name()))
			}
		}
	}

	return nil
}
