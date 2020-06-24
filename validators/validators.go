package validators

import (
	"strconv"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/tokenpool"
	"github.com/juju/errors"
)

var _ Factory = &Validators{}

type language interface {
	ID() uint16
	Code() string
}

// Factory represents a factory that generates new Validator types.
type Factory interface {
	NewForm(language language) (Validator, error)
	NewChoiceOption(value string, display string) ChoiceOption

	GetCountryOptionsForLanguage(language language) []ChoiceOption
	GetLanguageOptionsForLanguage(language language) []ChoiceOption
	GetCurrencyOptionsForLanguage(language language) []ChoiceOption
}

// Validators generates new Validator types.
type Validators struct {
	configService                config.Config
	loggerService                logger.Loggable
	languagesRepository          languages.Repository
	countriesRepository          countries.Repository
	currenciesRepository         currencies.Repository
	translationsRepository       translations.Repository
	regularExpressionsRepository regularexpressions.Repository
	tokenPoolService             tokenpool.Pool

	countryChoicesCache  *choicesCache
	languageChoicesCache *choicesCache
	currencyChoicesCache *choicesCache
}

// NewForm builds a new form-based validator based on the ValidatorFields given
func (validators *Validators) NewForm(language language) (Validator, error) {
	validator := &Form{
		configService:                validators.configService,
		loggerService:                validators.loggerService,
		translationsRepository:       validators.translationsRepository,
		tokenPoolService:             validators.tokenPoolService,
		regularExpressionsRepository: validators.regularExpressionsRepository,
		language:                     language,
		fields:                       map[string]*formField{},
		isFormValidator:              true,
	}

	if err := validator.AddField(validator.NewHoneypotField("_uname")); err != nil {
		return nil, errors.Trace(err)
	}
	token, err := validators.tokenPoolService.RequestToken()
	if err != nil {
		return validator, errors.Trace(err)
	}
	tokenField := validator.NewTokenField("_t")
	tokenField.SetValue(strconv.Itoa(token))
	if err := validator.AddField(tokenField); err != nil {
		return nil, errors.Trace(err)
	}

	return validator, nil
}

// New returns a new instance of Validators.
func New(configService config.Config, loggerService logger.Loggable, languagesRepository languages.Repository, countriesRepository countries.Repository, currenciesRepository currencies.Repository, translationsRepository translations.Repository, regularExpressionsRepository regularexpressions.Repository, tokenPoolService tokenpool.Pool) *Validators {
	return &Validators{
		configService:                configService,
		loggerService:                loggerService,
		languagesRepository:          languagesRepository,
		countriesRepository:          countriesRepository,
		currenciesRepository:         currenciesRepository,
		translationsRepository:       translationsRepository,
		regularExpressionsRepository: regularExpressionsRepository,
		tokenPoolService:             tokenPoolService,

		countryChoicesCache:  newChoicesCache(),
		languageChoicesCache: newChoicesCache(),
		currencyChoicesCache: newChoicesCache(),
	}
}
