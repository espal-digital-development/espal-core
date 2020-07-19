package runner

import (
	"io/ioutil"
	"path/filepath"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/image/optimizer"
	"github.com/espal-digital-development/espal-core/mailer"
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/espal-digital-development/espal-core/tokenpool"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/juju/errors"
)

func (r *Runner) core(path string) error {
	if path == "" {
		path = filepath.FromSlash("./app")
	}

	configYamlData, err := ioutil.ReadFile(filepath.FromSlash(path + "/config.yml"))
	if err != nil {
		return errors.Trace(err)
	}
	r.services.config, err = config.New(configYamlData)
	if err != nil {
		return errors.Trace(err)
	}

	if err := r.assets(); err != nil {
		return errors.Trace(err)
	}

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

	r.repositories.userRights = userrights.New()
	r.repositories.regularExpressions, err = regularexpressions.New()
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
	return nil
}
