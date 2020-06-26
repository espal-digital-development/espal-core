package runner

import (
	"path/filepath"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/mailer"
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/espal-digital-development/espal-core/storage/filesystem"
	"github.com/espal-digital-development/espal-core/tokenpool"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/juju/errors"
)

// nolint:funlen
func (r *Runner) core(path string) error {
	if path == "" {
		path = filepath.FromSlash("./app")
	}

	var err error
	r.storages.core, err = filesystem.New(path)
	if err != nil {
		return errors.Trace(err)
	}

	r.services.config, err = config.New(r.storages.core)
	if err != nil {
		return errors.Trace(err)
	}

	r.storages.translations, err = filesystem.New(r.services.config.TranslationsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsPrivateFiles, err = filesystem.New(r.services.config.PrivateFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsPublicFiles, err = filesystem.New(r.services.config.PublicFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsPublicRootFiles, err = filesystem.New(r.services.config.PublicRootFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsImages, err = filesystem.New(r.services.config.ImagesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsStylesheets, err = filesystem.New(r.services.config.StylesheetsAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	r.storages.assetsJavaScript, err = filesystem.New(r.services.config.JavaScriptAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}

	r.repositories.languages, err = languages.New()
	if err != nil {
		return errors.Trace(err)
	}
	r.services.logger = logger.New()

	availableLanguages := map[uint16]string{}
	for _, languageCode := range r.services.config.AvailableLanguages() {
		language, err := r.repositories.languages.ByCode(languageCode)
		if err != nil {
			return errors.Trace(err)
		}
		availableLanguages[language.ID()] = language.Code()
	}
	r.repositories.translations, err = translations.New(r.services.logger, r.storages.translations, availableLanguages)
	if err != nil {
		return errors.Trace(err)
	}
	r.repositories.countries, err = countries.New(r.repositories.languages, true)
	if err != nil {
		return errors.Trace(err)
	}
	r.repositories.currencies, err = currencies.New(r.repositories.languages)
	if err != nil {
		return errors.Trace(err)
	}

	r.repositories.userRights = userrights.New()
	r.repositories.regularExpressions, err = regularexpressions.New()
	if err != nil {
		return errors.Trace(err)
	}
	r.services.mailer = mailer.New(r.services.config.EmailHost(), r.services.config.EmailPort(),
		r.services.config.EmailUsername(), r.services.config.EmailPassword())
	if r.services.mailer == nil {
		return errors.Errorf("runner.services.mailer returned nil")
	}
	r.services.tokenPool = tokenpool.New(r.services.config.SecurityFormTokenLifespan(),
		r.services.config.SecurityFormTokenCleanupInterval())
	r.services.validators = validators.New(r.services.config, r.services.logger, r.repositories.languages,
		r.repositories.countries, r.repositories.currencies, r.repositories.translations,
		r.repositories.regularExpressions, r.services.tokenPool)
	return nil
}
