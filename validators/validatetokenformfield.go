package validators

import (
	"strconv"

	"github.com/juju/errors"
)

func (f *Form) validateTokenFormField(field FormField) (bool, error) {
	token, err := strconv.Atoi(field.Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !f.tokenPoolService.Validate(token) {
		field.AddError(f.translationsRepository.Singular(f.language.ID(), "validationTokenInvalidExpired"))

		token, err = f.tokenPoolService.RequestToken()
		if err != nil {
			return false, errors.Trace(err)
		}
		field.SetValue(strconv.Itoa(token))
	}
	return true, nil
}
