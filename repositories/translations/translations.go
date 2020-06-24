package translations

import (
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
	yaml "gopkg.in/yaml.v2"
)

var _ Repository = &Translations{}

const (
	errorNoTranslationWithForCode = "no translation language found for given code `%d`"
	errorNoTranslationKeyForName  = "no translation key found for given name `%s`"
	missingTranslationFormat      = "MISSING TRANSLATION (%s): %s\n"
)

// Repository represents a Translations repository.
type Repository interface {
	Singular(languageID uint16, key string) string
	Plural(languageID uint16, key string) string
	Formatted(languageID uint16, key string) string
	FormattedPlural(languageID uint16, key string) string
}

// Translations contains a full Translation repository.
type Translations struct {
	storage       storage.Storage
	loggerService logger.Loggable
	entries       map[uint16]map[string]translation
}

type translation struct {
	Singular          *string `yaml:"s"`
	Plural            *string `yaml:"p"`
	SingularFormatted *string `yaml:"sf"`
	PluralFormatted   *string `yaml:"pf"`
}

// Singular value based on languageCode.
func (translations *Translations) Singular(languageID uint16, key string) string {
	if !translations.checkExistence(languageID, key) {
		return key
	}
	if translations.entries[languageID][key].Singular == nil {
		translations.loggerService.Warningf(missingTranslationFormat, "Singular", key)
		return key
	}
	return *translations.entries[languageID][key].Singular
}

// Plural value based on languageID.
func (translations *Translations) Plural(languageID uint16, key string) string {
	if !translations.checkExistence(languageID, key) {
		return key
	}
	if translations.entries[languageID][key].Plural == nil {
		translations.loggerService.Warningf(missingTranslationFormat, "Plural", key)
		return key
	}
	return *translations.entries[languageID][key].Plural
}

// Formatted value based on languageID
func (translations *Translations) Formatted(languageID uint16, key string) string {
	if !translations.checkExistence(languageID, key) {
		return key
	}
	if translations.entries[languageID][key].SingularFormatted == nil {
		translations.loggerService.Warningf(missingTranslationFormat, "Plural", key)
		return key
	}
	return *translations.entries[languageID][key].SingularFormatted
}

// FormattedPlural value based on languageID
func (translations *Translations) FormattedPlural(languageID uint16, key string) string {
	if !translations.checkExistence(languageID, key) {
		return key
	}
	if translations.entries[languageID][key].PluralFormatted == nil {
		translations.loggerService.Warningf(missingTranslationFormat, "PluralFormatted", key)
		return key
	}
	return *translations.entries[languageID][key].PluralFormatted
}

func (translations *Translations) checkExistence(languageID uint16, key string) bool {
	if _, ok := translations.entries[languageID]; !ok {
		translations.loggerService.Warningf(errorNoTranslationWithForCode, languageID)
		return false
	}
	if _, ok := translations.entries[languageID][key]; !ok {
		translations.loggerService.Warningf(errorNoTranslationKeyForName, key)
		return false
	}
	return true
}

// New returns new a Languages repository instance.
func New(loggerService logger.Loggable, storage storage.Storage, availableLanguages map[uint16]string) (*Translations, error) {
	translations := &Translations{
		loggerService: loggerService,
		storage:       storage,
		entries:       map[uint16]map[string]translation{},
	}

	for languageID, language := range availableLanguages {
		file, ok, err := storage.Get(language + ".yml")
		if err != nil {
			return nil, errors.Trace(err)
		}
		if !ok {
			continue
		}
		translations.entries[languageID] = map[string]translation{}
		err = yaml.Unmarshal(file, translations.entries[languageID])
		if err != nil {
			return nil, errors.Trace(err)
		}
	}

	return translations, nil
}
