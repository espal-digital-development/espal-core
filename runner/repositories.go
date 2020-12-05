package runner

import (
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/juju/errors"
)

type repositories struct {
	regularExpressions regularexpressions.Repository
	userRights         userrights.Repository
	languages          languages.Repository
	translations       translations.Repository
	countries          countries.Repository
	currencies         currencies.Repository
	themes             themes.Repository
}

func (r *Runner) repos() error {
	for k := range r.modulesRegistry {
		moduleRepos, err := r.modulesRegistry[k].GetRepositories()
		if err != nil {
			return errors.Trace(err)
		}
		moduleRepos.SetRegularExpressions(r.repositories.regularExpressions)
		moduleRepos.SetUserRights(r.repositories.userRights)
		moduleRepos.SetLanguages(r.repositories.languages)
		moduleRepos.SetTranslations(r.repositories.translations)
		moduleRepos.SetCountries(r.repositories.countries)
		moduleRepos.SetCurrencies(r.repositories.currencies)
		moduleRepos.SetThemes(r.repositories.themes)
	}
	return nil
}
