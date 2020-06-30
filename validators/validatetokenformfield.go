package validators

import (
	"bytes"
	"strconv"

	"github.com/espal-digital-development/espal-core/sessions"
	"github.com/juju/errors"
)

func (f *Form) validateTokenFormField(field FormField) (bool, error) {
	tokenInSession, ok, err := f.context.GetSessionValue(sessions.SessionKeyFormToken)
	if err != nil {
		return false, errors.Trace(err)
	}

	tokenIsValid := ok && bytes.Equal(tokenInSession, field.ValueAsBytes())
	if tokenIsValid {
		token, err := strconv.Atoi(field.Value())
		if err != nil {
			return false, errors.Trace(err)
		}
		tokenIsValid = f.tokenPoolService.Validate(token)
	}

	if !tokenIsValid {
		field.AddError(f.translationsRepository.Singular(f.language.ID(), "validationTokenInvalidExpired"))

		token, err := f.tokenPoolService.RequestToken()
		if err != nil {
			return false, errors.Trace(err)
		}
		field.SetValue(strconv.Itoa(token))
	}
	return true, nil
}
