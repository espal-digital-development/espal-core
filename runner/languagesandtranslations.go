package runner

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/juju/errors"
)

const notAvailableLanguageErrTemplate = "cannot load translations from module `%s` for language " +
	"`%s` as it's not an available language"

func (r *Runner) languagesAndTranslations() error {
	availableLanguages := map[uint16]string{}
	for _, languageCode := range r.services.config.AvailableLanguages() {
		language, err := r.repositories.languages.ByCode(languageCode)
		if err != nil {
			return errors.Trace(err)
		}
		availableLanguages[language.ID()] = language.Code()
	}

	translations, err := translations.New(r.services.logger, r.storages.translations, availableLanguages)
	if err != nil {
		return errors.Trace(err)
	}
	for k := range r.modulesRegistry {
		moduleTranslations, err := r.modulesRegistry[k].GetTranslations()
		if err != nil {
			return errors.Trace(err)
		}
		if moduleTranslations == nil {
			continue
		}
		data, err := moduleTranslations.GetAll()
		if err != nil {
			return errors.Trace(err)
		}
		for language, filePath := range data {
			var languageID uint16
			for availableLanguageID, languageCode := range availableLanguages {
				if languageCode == language {
					languageID = availableLanguageID
				}
			}
			if languageID == 0 {
				return errors.Errorf(
					notAvailableLanguageErrTemplate,
					r.modulesRegistry[k].GetMeta().Name(), language,
				)
			}
			if err := translations.LoadForLanguageFromYaml(languageID, filePath); err != nil {
				return errors.Trace(err)
			}
		}
	}
	translationsPath := r.services.config.TranslationsPath()
	if translationsPath != "" {
		files, err := ioutil.ReadDir(translationsPath)
		if err != nil {
			return errors.Trace(err)
		}
		for k := range files {
			if files[k].IsDir() {
				continue
			}
			ext := filepath.Ext(files[k].Name())
			if ext != ".yml" {
				continue
			}
			language := strings.TrimSuffix(files[k].Name(), ext)
			var languageID uint16
			for availableLanguageID, languageCode := range availableLanguages {
				if languageCode == language {
					languageID = availableLanguageID
				}
			}
			if languageID == 0 {
				return errors.Errorf(
					notAvailableLanguageErrTemplate,
					r.modulesRegistry[k].GetMeta().Name(), language,
				)
			}
			if err := translations.LoadForLanguageFromYaml(languageID,
				filepath.FromSlash(translationsPath+"/"+files[k].Name())); err != nil {
				return errors.Trace(err)
			}
		}
	}
	r.repositories.translations = translations

	return nil
}
