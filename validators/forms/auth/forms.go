package auth

import (
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Factory = &Forms{}

type language interface {
	ID() uint16
	Code() string
}

// Factory represents an object that serves new forms.
type Factory interface {
	New(language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory validators.Factory
	userStore         user.Store
}

// New creates a new Form instance with the required logic.
func (forms *Forms) New(language language) (Form, error) {
	validator, err := forms.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	emailField := validator.NewEmailField("email")
	emailField.SetValidate()
	emailField.SetHideLabel()
	if err := validator.AddField(emailField); err != nil {
		return nil, errors.Trace(err)
	}

	passwordField := validator.NewPasswordField("password")
	passwordField.SetHideLabel()
	if err := validator.AddField(passwordField); err != nil {
		return nil, errors.Trace(err)
	}

	rememberMeField := validator.NewCheckboxField("rememberMe")
	rememberMeField.SetHideLabel()
	if err := validator.AddField(rememberMeField); err != nil {
		return nil, errors.Trace(err)
	}

	return &Auth{
		validator: validator,
		userStore: forms.userStore,
		view:      formview.New(validator),
	}, nil
}

// New returns a new instance of LoginForm.
func New(validatorsFactory validators.Factory, userStore user.Store) *Forms {
	return &Forms{
		validatorsFactory: validatorsFactory,
		userStore:         userStore,
	}
}
