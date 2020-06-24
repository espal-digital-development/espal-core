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

func (runner *Runner) core(path string) error {
	if path == "" {
		path = filepath.FromSlash("./app")
	}

	var err error
	runner.storages.core, err = filesystem.New(path)
	if err != nil {
		return errors.Trace(err)
	}

	runner.services.config, err = config.New(runner.storages.core)
	if err != nil {
		return errors.Trace(err)
	}

	runner.storages.translations, err = filesystem.New(runner.services.config.TranslationsPath())
	if err != nil {
		return errors.Trace(err)
	}
	runner.storages.assetsPrivateFiles, err = filesystem.New(runner.services.config.PrivateFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	runner.storages.assetsPublicFiles, err = filesystem.New(runner.services.config.PublicFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	runner.storages.assetsPublicRootFiles, err = filesystem.New(runner.services.config.PublicRootFilesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	runner.storages.assetsImages, err = filesystem.New(runner.services.config.ImagesAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	runner.storages.assetsStylesheets, err = filesystem.New(runner.services.config.StylesheetsAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}
	runner.storages.assetsJavaScript, err = filesystem.New(runner.services.config.JavaScriptAssetsPath())
	if err != nil {
		return errors.Trace(err)
	}

	runner.repositories.languages, err = languages.New()
	if err != nil {
		return errors.Trace(err)
	}
	runner.services.logger = logger.New()

	availableLanguages := map[uint16]string{}
	for _, languageCode := range runner.services.config.AvailableLanguages() {
		language, err := runner.repositories.languages.ByCode(languageCode)
		if err != nil {
			return errors.Trace(err)
		}
		availableLanguages[language.ID()] = language.Code()
	}
	runner.repositories.translations, err = translations.New(runner.services.logger, runner.storages.translations, availableLanguages)
	if err != nil {
		return errors.Trace(err)
	}
	runner.repositories.countries, err = countries.New(runner.repositories.languages, true)
	if err != nil {
		return errors.Trace(err)
	}
	runner.repositories.currencies, err = currencies.New(runner.repositories.languages)
	if err != nil {
		return errors.Trace(err)
	}

	runner.repositories.userRights = userrights.New()
	runner.repositories.regularExpressions, err = regularexpressions.New()
	if err != nil {
		return errors.Trace(err)
	}
	runner.services.mailer = mailer.New(runner.services.config.EmailHost(), runner.services.config.EmailPort(), runner.services.config.EmailUsername(), runner.services.config.EmailPassword())
	if runner.services.mailer == nil {
		return errors.Errorf("runner.services.mailer returned nil")
	}
	runner.services.tokenPool = tokenpool.New(runner.services.config.SecurityFormTokenLifespan(), runner.services.config.SecurityFormTokenCleanupInterval())
	runner.services.validators = validators.New(runner.services.config, runner.services.logger, runner.repositories.languages, runner.repositories.countries, runner.repositories.currencies, runner.repositories.translations, runner.repositories.regularExpressions, runner.services.tokenPool)
	return nil
}
