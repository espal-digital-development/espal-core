package renderer

import (
	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/juju/errors"
)

var _ Renderer = &TemplateRenderer{}

type context interface {
	Translate(string) string
	TranslatePlural(string) string
	AdminURL() string
}

type entity interface {
	ID() string
	IsUpdated() bool
	CreatedByID() string
	UpdatedByID() *string
	CreatedByFirstName() *string
	CreatedBySurname() *string
	UpdatedByFirstName() *string
	UpdatedBySurname() *string
}

// Renderer represents an object that offers interactions used in
// the rendering of visual presentation layers.
type Renderer interface {
	CreatedBy(entity entity, languageID uint16) string
	UpdatedBy(entity entity, languageID uint16) string
	CountryName(countryID uint16, languageID uint16) string
	LanguageName(languageID uint16, targetLanguageID uint16) string
	DefaultOverviewTranslations(ctx context) string
	CreatedUpdatedByLinks(ctx context, languageID uint16, entity entity) string
	AccountMenu(ctx context) string
}

// TemplateRenderer templating facility service object for rendering simple
// parts of printable output.
type TemplateRenderer struct {
	languagesRepository    languages.Repository
	countriesRepository    countries.Repository
	translationsRepository translations.Repository
	loggerService          logger.Loggable
}

// CreatedBy returns the presentable name for the User that created this entity.
func (r *TemplateRenderer) CreatedBy(entity entity, languageID uint16) string {
	var createdBy string
	if entity.CreatedByFirstName() != nil {
		createdBy = *entity.CreatedByFirstName()
	}
	if entity.CreatedBySurname() != nil {
		createdBy += " " + *entity.CreatedBySurname()
	}
	if createdBy == "" {
		createdBy = r.translationsRepository.Singular(languageID, "user") + " " + entity.CreatedByID()
	}
	return createdBy
}

// UpdatedBy returns the presentable name for the User that last updated this entity.
func (r *TemplateRenderer) UpdatedBy(entity entity, languageID uint16) string {
	if entity.UpdatedByID() == nil {
		return ""
	}
	var updatedBy string
	if entity.UpdatedByFirstName() != nil {
		updatedBy = *entity.UpdatedByFirstName()
	}
	if entity.UpdatedBySurname() != nil {
		updatedBy += " " + *entity.UpdatedBySurname()
	}
	if updatedBy == "" {
		updatedBy = r.translationsRepository.Singular(languageID, "user") + " " + *entity.UpdatedByID()
	}
	return updatedBy
}

// CountryName returns the localized name for the given countryID and languageID.
func (r *TemplateRenderer) CountryName(countryID uint16, languageID uint16) string {
	country, err := r.countriesRepository.ByID(countryID)
	if err != nil {
		r.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	return country.Translate(languageID)
}

// LanguageName returns the localized name for the given languageID and targetLanguageID.
func (r *TemplateRenderer) LanguageName(languageID uint16, targetLanguageID uint16) string {
	language, err := r.languagesRepository.ByID(languageID)
	if err != nil {
		r.loggerService.Error(errors.ErrorStack(err))
		return ""
	}
	return language.Translate(targetLanguageID)
}

func (r *TemplateRenderer) perror(i int, err error) {
	if err != nil {
		r.loggerService.Error(errors.ErrorStack(err))
	}
}

// New returns a new instance of Renderer.
func New(languagesRepository languages.Repository, countriesRepository countries.Repository, translationsRepository translations.Repository, loggerService logger.Loggable) *TemplateRenderer {
	return &TemplateRenderer{
		languagesRepository:    languagesRepository,
		countriesRepository:    countriesRepository,
		translationsRepository: translationsRepository,
		loggerService:          loggerService,
	}
}
