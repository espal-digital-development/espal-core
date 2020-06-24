package validators

import (
	"strconv"

	"github.com/juju/errors"
)

func (form *Form) validateTokenFormField(field FormField) (bool, error) {
	token, err := strconv.Atoi(field.Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !form.tokenPoolService.Validate(token) {
		field.AddError(form.translationsRepository.Singular(form.language.ID(), "validationTokenInvalidExpired"))

		token, err = form.tokenPoolService.RequestToken()
		if err != nil {
			return false, errors.Trace(err)
		}
		field.SetValue(strconv.Itoa(token))
	}
	return true, nil
}
