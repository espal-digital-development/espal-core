package translations

import (
	"strings"

	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/translations/translationsdata"
	"github.com/juju/errors"
	yaml "gopkg.in/yaml.v2"
)

var _ Repository = &Translations{}

const (
	errorNoTranslationWithForCode = "no translation language found for given code `%d`"
	errorNoTranslationKeyForName  = "no translation key found for given name `%s`"
	missingTranslationFormat      = "MISSING TRANSLATION (%s): %s\n"
	errAlreadySetForLanguageID    = "%s translation already set for key `%s`"
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
	loggerService logger.Loggable
	entries       map[uint16]map[string]translation
}

type translation struct {
	Singular          *string `yaml:"s"`
	Plural            *string `yaml:"p"`
	SingularFormatted *string `yaml:"sf"`
	PluralFormatted   *string `yaml:"pf"`
}

// SetSingular sets the singular value based on languageCode and key.
func (t *Translations) SetSingular(languageID uint16, key string, value string) error {
	if t.checkExistence(languageID, key) {
		return errors.Errorf(errAlreadySetForLanguageID, "singular", key)
	}
	if _, ok := t.entries[languageID]; !ok {
		t.entries[languageID] = map[string]translation{}
	}
	if _, ok := t.entries[languageID][key]; !ok {
		t.entries[languageID][key] = translation{
			Singular: &value,
		}
	}
	*t.entries[languageID][key].Singular = value
	return nil
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

// SetPlural sets the plural value based on languageCode and key.
func (t *Translations) SetPlural(languageID uint16, key string, value string) error {
	if t.checkExistence(languageID, key) {
		return errors.Errorf(errAlreadySetForLanguageID, "plural", key)
	}
	if _, ok := t.entries[languageID]; !ok {
		t.entries[languageID] = map[string]translation{}
	}
	if _, ok := t.entries[languageID][key]; !ok {
		t.entries[languageID][key] = translation{
			Plural: &value,
		}
	}
	*t.entries[languageID][key].Plural = value
	return nil
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

// SetFormatted sets the formatted value based on languageCode and key.
func (t *Translations) SetFormatted(languageID uint16, key string, value string) error {
	if t.checkExistence(languageID, key) {
		return errors.Errorf(errAlreadySetForLanguageID, "formatted", key)
	}
	if _, ok := t.entries[languageID]; !ok {
		t.entries[languageID] = map[string]translation{}
	}
	if _, ok := t.entries[languageID][key]; !ok {
		t.entries[languageID][key] = translation{
			SingularFormatted: &value,
		}
	}
	*t.entries[languageID][key].SingularFormatted = value
	return nil
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

// SetFormattedPlural sets the plural formatted value based on languageCode and key.
func (t *Translations) SetFormattedPlural(languageID uint16, key string, value string) error {
	if t.checkExistence(languageID, key) {
		return errors.Errorf(errAlreadySetForLanguageID, "pluralFormatted", key)
	}
	if _, ok := t.entries[languageID]; !ok {
		t.entries[languageID] = map[string]translation{}
	}
	if _, ok := t.entries[languageID][key]; !ok {
		t.entries[languageID][key] = translation{
			PluralFormatted: &value,
		}
	}
	*t.entries[languageID][key].PluralFormatted = value
	return nil
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

// LoadForLanguageFromYamlData loads in the given yaml translation data from the given path for the given language.
func (t *Translations) LoadForLanguageFromYamlData(languageID uint16, yamlData []byte) error {
	data := map[string]translation{}
	if err := yaml.Unmarshal(yamlData, data); err != nil {
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
	for key, trans := range data {
		if _, ok := t.entries[languageID]; !ok {
			t.entries[languageID] = map[string]translation{}
		}
		if _, ok := t.entries[languageID][key]; !ok {
			t.entries[languageID][key] = trans
			continue
		}

		entry := t.entries[languageID][key]
		if trans.Plural != nil && *trans.Plural != "" {
			entry.Plural = trans.Plural
		}
		if trans.PluralFormatted != nil && *trans.PluralFormatted != "" {
			entry.PluralFormatted = trans.PluralFormatted
		}
		if trans.Singular != nil && *trans.Singular != "" {
			entry.Singular = trans.Singular
		}
		if trans.SingularFormatted != nil && *trans.SingularFormatted != "" {
			entry.SingularFormatted = trans.SingularFormatted
		}
	}
}

// New returns new a Languages repository instance.
func New(loggerService logger.Loggable, availableLanguages map[uint16]string) (*Translations, error) {
	t := &Translations{
		loggerService: loggerService,
		entries:       map[uint16]map[string]translation{},
	}

	for languageID, language := range availableLanguages {
		coreFile, err := translationsdata.Asset("_data/" + language + ".yml")
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
