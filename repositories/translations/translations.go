package translations

import (
	"io/ioutil"
	"strings"

	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/translations/translationsdata"
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
func (t *Translations) Singular(languageID uint16, key string) string {
	if !t.checkExistence(languageID, key) {
		return key
	}
	if t.entries[languageID][key].Singular == nil {
		t.loggerService.Warningf(missingTranslationFormat, "Singular", key)
		return key
	}
	return *t.entries[languageID][key].Singular
}

// Plural value based on languageID.
func (t *Translations) Plural(languageID uint16, key string) string {
	if !t.checkExistence(languageID, key) {
		return key
	}
	if t.entries[languageID][key].Plural == nil {
		t.loggerService.Warningf(missingTranslationFormat, "Plural", key)
		return key
	}
	return *t.entries[languageID][key].Plural
}

// Formatted value based on languageID.
func (t *Translations) Formatted(languageID uint16, key string) string {
	if !t.checkExistence(languageID, key) {
		return key
	}
	if t.entries[languageID][key].SingularFormatted == nil {
		t.loggerService.Warningf(missingTranslationFormat, "Plural", key)
		return key
	}
	return *t.entries[languageID][key].SingularFormatted
}

// FormattedPlural value based on languageID.
func (t *Translations) FormattedPlural(languageID uint16, key string) string {
	if !t.checkExistence(languageID, key) {
		return key
	}
	if t.entries[languageID][key].PluralFormatted == nil {
		t.loggerService.Warningf(missingTranslationFormat, "PluralFormatted", key)
		return key
	}
	return *t.entries[languageID][key].PluralFormatted
}

// LoadForLanguageFromYaml loads in the given yaml translation data from the given path for the given language.
func (t *Translations) LoadForLanguageFromYaml(languageID uint16, path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Trace(err)
	}
	data := map[string]translation{}
	err = yaml.Unmarshal(file, data)
	if err != nil {
		return errors.Trace(err)
	}
	t.loadFromData(languageID, data)
	return nil
}

func (t *Translations) checkExistence(languageID uint16, key string) bool {
	if _, ok := t.entries[languageID]; !ok {
		t.loggerService.Warningf(errorNoTranslationWithForCode, languageID)
		return false
	}
	if _, ok := t.entries[languageID][key]; !ok {
		t.loggerService.Warningf(errorNoTranslationKeyForName, key)
		return false
	}
	return true
}

func (t *Translations) loadFromData(languageID uint16, data map[string]translation) {
	for key, translation := range data {
		if _, ok := t.entries[languageID][key]; !ok {
			t.entries[languageID][key] = translation
			continue
		}

		if translation.Plural != nil && *translation.Plural != "" {
			*t.entries[languageID][key].Plural = *translation.Plural
		}
		if translation.PluralFormatted != nil && *translation.PluralFormatted != "" {
			*t.entries[languageID][key].PluralFormatted = *translation.PluralFormatted
		}
		if translation.Singular != nil && *translation.Singular != "" {
			*t.entries[languageID][key].Singular = *translation.Singular
		}
		if translation.SingularFormatted != nil && *translation.SingularFormatted != "" {
			*t.entries[languageID][key].SingularFormatted = *translation.SingularFormatted
		}
	}
}

// New returns new a Languages repository instance.
func New(loggerService logger.Loggable, storage storage.Storage,
	availableLanguages map[uint16]string) (*Translations, error) {
	t := &Translations{
		loggerService: loggerService,
		storage:       storage,
		entries:       map[uint16]map[string]translation{},
	}

	for languageID, language := range availableLanguages {
		coreFile, err := translationsdata.Asset("_data/" + language + ".yml")
		if err != nil && !strings.HasSuffix(err.Error(), " not found") {
			return nil, errors.Trace(err)
		}
		if err == nil || !strings.HasSuffix(err.Error(), " not found") {
			t.entries[languageID] = map[string]translation{}
			err = yaml.Unmarshal(coreFile, t.entries[languageID])
			if err != nil {
				return nil, errors.Trace(err)
			}
		}
	}

	return t, nil
}
