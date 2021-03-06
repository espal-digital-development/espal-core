// Code generated by espal-store-synthesizer. DO NOT EDIT.
package domain

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ DomainEntity = &Domain{}

type DomainEntity interface {
	database.Model
	SiteID() string
	SetSiteID(siteID string)
	Active() bool
	SetActive(active bool)
	Host() string
	SetHost(host string)
	Language() *uint16
	SetLanguage(language *uint16)
	Currencies() string
	SetCurrencies(currencies string)
	LocalizedName() *string
	SetLocalizedName(localizedName *string)
	CurrenciesCount() uint
	HostWithProtocol() string
	HostWithProtocolAndWWW() string
}

// TableName returns the table name that belongs to the current model.
func (d *Domain) TableName() string {
	return "Domain"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (d *Domain) TableAlias() string {
	return "de"
}

// ID returns id.
func (d *Domain) ID() string {
	return d.id
}

// CreatedByID returns createdByID.
func (d *Domain) CreatedByID() string {
	return d.createdByID
}

// SetCreatedByID sets the createdByID.
func (d *Domain) SetCreatedByID(createdByID string) {
	d.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (d *Domain) UpdatedByID() *string {
	return d.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (d *Domain) SetUpdatedByID(updatedByID *string) {
	d.updatedByID = updatedByID
}

// SiteID returns siteID.
func (d *Domain) SiteID() string {
	return d.siteID
}

// SetSiteID sets the siteID.
func (d *Domain) SetSiteID(siteID string) {
	d.siteID = siteID
}

// CreatedAt returns createdAt.
func (d *Domain) CreatedAt() time.Time {
	return d.createdAt
}

// SetCreatedAt sets the createdAt.
func (d *Domain) SetCreatedAt(createdAt time.Time) {
	d.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (d *Domain) UpdatedAt() *time.Time {
	return d.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (d *Domain) SetUpdatedAt(updatedAt *time.Time) {
	d.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (d *Domain) CreatedByFirstName() *string {
	return d.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (d *Domain) SetCreatedByFirstName(createdByFirstName *string) {
	d.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (d *Domain) CreatedBySurname() *string {
	return d.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (d *Domain) SetCreatedBySurname(createdBySurname *string) {
	d.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (d *Domain) UpdatedByFirstName() *string {
	return d.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (d *Domain) SetUpdatedByFirstName(updatedByFirstName *string) {
	d.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (d *Domain) UpdatedBySurname() *string {
	return d.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (d *Domain) SetUpdatedBySurname(updatedBySurname *string) {
	d.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (d *Domain) IsUpdated() bool {
	return d.updatedByID != nil
}

// Active returns active.
func (d *Domain) Active() bool {
	return d.active
}

// SetActive sets the active.
func (d *Domain) SetActive(active bool) {
	d.active = active
}

// Host returns host.
func (d *Domain) Host() string {
	return d.host
}

// SetHost sets the host.
func (d *Domain) SetHost(host string) {
	d.host = host
}

// Language returns language.
func (d *Domain) Language() *uint16 {
	return d.language
}

// SetLanguage sets the language.
func (d *Domain) SetLanguage(language *uint16) {
	d.language = language
}

// Currencies returns currencies.
func (d *Domain) Currencies() string {
	return d.currencies
}

// SetCurrencies sets the currencies.
func (d *Domain) SetCurrencies(currencies string) {
	d.currencies = currencies
}

// LocalizedName returns localizedName.
func (d *Domain) LocalizedName() *string {
	return d.localizedName
}

// SetLocalizedName sets the localizedName.
func (d *Domain) SetLocalizedName(localizedName *string) {
	d.localizedName = localizedName
}

func newDomain() *Domain {
	return &Domain{}
}

// New returns a new instance of DomainEntity.
func NewDomainEntity() DomainEntity {
	return newDomain()
}
