package runner

import (
	"io/ioutil"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/image/optimizer"
	"github.com/espal-digital-development/espal-core/mailer"
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/espal-digital-development/espal-core/tokenpool"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/juju/errors"
)

func (r *Runner) core(path string) error {
	if path == "" {
		path = "./app"
	}

	configYamlData, err := ioutil.ReadFile(path + "/config.yml")
	if err != nil {
		return errors.Trace(err)
	}
	r.services.config, err = config.New(configYamlData)
	if err != nil {
		return errors.Trace(err)
	}

	for k := range r.modulesRegistry {
		moduleConfig, err := r.modulesRegistry[k].GetConfig()
		if err != nil {
			return errors.Trace(err)
		}
		moduleConfig.SetService(r.services.config)
	}

	if err := r.assets(); err != nil {
		return errors.Trace(err)
	}

	r.repositories.regularExpressions, err = regularexpressions.New()
	if err != nil {
		return errors.Trace(err)
	}
	r.repositories.userRights = userrights.New()

	r.repositories.languages, err = languages.New()
	if err != nil {
		return errors.Trace(err)
	}

	if err := r.languagesAndTranslations(); err != nil {
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

	r.repositories.themes, err = themes.New()
	if err != nil {
		return errors.Trace(err)
	}

	r.services.imageOptimizer, err = optimizer.New(r.services.config)
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
	for k := range r.modulesRegistry {
		r.modulesRegistry[k].RegisterValidatorsFactory(r.services.validators)
	}
	return nil
}
