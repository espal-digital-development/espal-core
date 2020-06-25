// Code generated by espal-store-synthesizer. DO NOT EDIT.
package site

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ SiteEntity = &Site{}

type SiteEntity interface {
	database.Model
	Online() bool
	SetOnline(online bool)
	Language() *uint16
	SetLanguage(language *uint16)
	Country() *uint16
	SetCountry(country *uint16)
	Currencies() string
	SetCurrencies(currencies string)
	LocalizedName() *string
	SetLocalizedName(localizedName *string)
	CurrenciesCount() uint
}

// TableName returns the table name that belongs to the current model.
func (s *Site) TableName() string {
	return "Site"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *Site) TableAlias() string {
	return "se"
}

// ID returns id.
func (s *Site) ID() string {
	return s.id
}

// CreatedByID returns createdByID.
func (s *Site) CreatedByID() string {
	return s.createdByID
}

// SetCreatedByID sets the createdByID.
func (s *Site) SetCreatedByID(createdByID string) {
	s.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (s *Site) UpdatedByID() *string {
	return s.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (s *Site) SetUpdatedByID(updatedByID *string) {
	s.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (s *Site) CreatedAt() time.Time {
	return s.createdAt
}

// SetCreatedAt sets the createdAt.
func (s *Site) SetCreatedAt(createdAt time.Time) {
	s.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (s *Site) UpdatedAt() *time.Time {
	return s.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (s *Site) SetUpdatedAt(updatedAt *time.Time) {
	s.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (s *Site) CreatedByFirstName() *string {
	return s.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (s *Site) SetCreatedByFirstName(createdByFirstName *string) {
	s.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (s *Site) CreatedBySurname() *string {
	return s.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (s *Site) SetCreatedBySurname(createdBySurname *string) {
	s.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (s *Site) UpdatedByFirstName() *string {
	return s.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (s *Site) SetUpdatedByFirstName(updatedByFirstName *string) {
	s.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (s *Site) UpdatedBySurname() *string {
	return s.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (s *Site) SetUpdatedBySurname(updatedBySurname *string) {
	s.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (s *Site) IsUpdated() bool {
	return s.updatedByID != nil
}

// Online returns online.
func (s *Site) Online() bool {
	return s.online
}

// SetOnline sets the online.
func (s *Site) SetOnline(online bool) {
	s.online = online
}

// Language returns language.
func (s *Site) Language() *uint16 {
	return s.language
}

// SetLanguage sets the language.
func (s *Site) SetLanguage(language *uint16) {
	s.language = language
}

// Country returns country.
func (s *Site) Country() *uint16 {
	return s.country
}

// SetCountry sets the country.
func (s *Site) SetCountry(country *uint16) {
	s.country = country
}

// Currencies returns currencies.
func (s *Site) Currencies() string {
	return s.currencies
}

// SetCurrencies sets the currencies.
func (s *Site) SetCurrencies(currencies string) {
	s.currencies = currencies
}

// LocalizedName returns localizedName.
func (s *Site) LocalizedName() *string {
	return s.localizedName
}

// SetLocalizedName sets the localizedName.
func (s *Site) SetLocalizedName(localizedName *string) {
	s.localizedName = localizedName
}

func newSite() *Site {
	return &Site{}
}

// New returns a new instance of SiteEntity.
func NewSiteEntity() SiteEntity {
	return newSite()
}
