// Code generated by espal-store-synthesizer. DO NOT EDIT.
package group

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ GroupEntity = &Group{}

type GroupEntity interface {
	database.Model
	Active() bool
	SetActive(active bool)
	UserRights() string
	SetUserRights(userRights string)
	Currencies() string
	SetCurrencies(currencies string)
	LocalizedName() *string
	SetLocalizedName(localizedName *string)
	UserRightsCount() uint
	CurrenciesCount() uint
}

// ID returns id.
func (g *Group) ID() string {
	return g.id
}

// CreatedByID returns createdByID.
func (g *Group) CreatedByID() string {
	return g.createdByID
}

// SetCreatedByID sets the createdByID.
func (g *Group) SetCreatedByID(createdByID string) {
	g.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (g *Group) UpdatedByID() *string {
	return g.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (g *Group) SetUpdatedByID(updatedByID *string) {
	g.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (g *Group) CreatedAt() time.Time {
	return g.createdAt
}

// SetCreatedAt sets the createdAt.
func (g *Group) SetCreatedAt(createdAt time.Time) {
	g.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (g *Group) UpdatedAt() *time.Time {
	return g.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (g *Group) SetUpdatedAt(updatedAt *time.Time) {
	g.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (g *Group) CreatedByFirstName() *string {
	return g.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (g *Group) SetCreatedByFirstName(createdByFirstName *string) {
	g.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (g *Group) CreatedBySurname() *string {
	return g.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (g *Group) SetCreatedBySurname(createdBySurname *string) {
	g.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (g *Group) UpdatedByFirstName() *string {
	return g.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (g *Group) SetUpdatedByFirstName(updatedByFirstName *string) {
	g.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (g *Group) UpdatedBySurname() *string {
	return g.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (g *Group) SetUpdatedBySurname(updatedBySurname *string) {
	g.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (g *Group) IsUpdated() bool {
	return g.updatedByID != nil
}

// Active returns active.
func (g *Group) Active() bool {
	return g.active
}

// SetActive sets the active.
func (g *Group) SetActive(active bool) {
	g.active = active
}

// UserRights returns userRights.
func (g *Group) UserRights() string {
	return g.userRights
}

// SetUserRights sets the userRights.
func (g *Group) SetUserRights(userRights string) {
	g.userRights = userRights
}

// Currencies returns currencies.
func (g *Group) Currencies() string {
	return g.currencies
}

// SetCurrencies sets the currencies.
func (g *Group) SetCurrencies(currencies string) {
	g.currencies = currencies
}

// LocalizedName returns localizedName.
func (g *Group) LocalizedName() *string {
	return g.localizedName
}

// SetLocalizedName sets the localizedName.
func (g *Group) SetLocalizedName(localizedName *string) {
	g.localizedName = localizedName
}

func newGroup() *Group {
	return &Group{}
}

// New returns a new instance of GroupEntity.
func NewGroupEntity() GroupEntity {
	return newGroup()
}
