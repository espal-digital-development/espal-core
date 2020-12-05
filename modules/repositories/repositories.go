package repositories

import (
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
)

type Repositories struct {
	regularExpressions regularexpressions.Repository
	userRights         userrights.Repository
	languages          languages.Repository
	translations       translations.Repository
	countries          countries.Repository
	currencies         currencies.Repository
	themes             themes.Repository
}

// SetRegularExpressions sets the RegularExpressions repository.
func (r *Repositories) SetRegularExpressions(regularExpressions regularexpressions.Repository) {
	r.regularExpressions = regularExpressions
}

// RegularExpressions returns the RegularExpressions repository.
func (r *Repositories) RegularExpressions() regularexpressions.Repository {
	return r.regularExpressions
}

// UserRights sets the UserRights repository.
func (r *Repositories) SetUserRights(userRights userrights.Repository) {
	r.userRights = userRights
}

// UserRights returns the UserRights repository.
func (r *Repositories) UserRights() userrights.Repository {
	return r.userRights
}

// Languages sets the Languages repository.
func (r *Repositories) SetLanguages(languages languages.Repository) {
	r.languages = languages
}

// Languages returns the Languages repository.
func (r *Repositories) Languages() languages.Repository {
	return r.languages
}

// Translations sets the Translations repository.
func (r *Repositories) SetTranslations(translations translations.Repository) {
	r.translations = translations
}

// Translations returns the Translations repository.
func (r *Repositories) Translations() translations.Repository {
	return r.translations
}

// Countries sets the Countries repository.
func (r *Repositories) SetCountries(countries countries.Repository) {
	r.countries = countries
}

// Countries returns the Countries repository.
func (r *Repositories) Countries() countries.Repository {
	return r.countries
}

// Currencies sets the Currencies repository.
func (r *Repositories) SetCurrencies(currencies currencies.Repository) {
	r.currencies = currencies
}

// Currencies returns the Currencies repository.
func (r *Repositories) Currencies() currencies.Repository {
	return r.currencies
}

// Themes sets the Themes repository.
func (r *Repositories) SetThemes(themes themes.Repository) {
	r.themes = themes
}

// Themes returns the Themes repository.
func (r *Repositories) Themes() themes.Repository {
	return r.themes
}

// New returns a new instance of Repositories.
func New() (*Repositories, error) {
	r := &Repositories{}
	return r, nil
}
