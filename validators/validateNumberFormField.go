package validators

import (
	"fmt"
	"strconv"

	"github.com/juju/errors"
)

func (f *Form) validateNumberFormField(field *formField) error {
	var err error
	var value float64
	if len(field.Value()) > 0 {
		value, err = strconv.ParseFloat(field.Value(), 64)
		if err != nil {
			return errors.Trace(err)
		}
	}

	if field.Optional() && value == 0 {
		// Empty fields allowed on `optional` and having a `minLength`
	} else if field.MinValue() > 0 && value < field.MinValue() {
		field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(), "fieldXCannotBeSmallerThanX"), field.Name(), field.MinValue()))
	}
	if field.MaxValue() > 0 && value > field.MaxValue() {
		field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(), "fieldXCannotBeGreaterThanX"), field.Name(), field.MaxValue()))
	}

	return nil
}
