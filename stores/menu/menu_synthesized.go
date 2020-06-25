// Code generated by espal-store-synthesizer. DO NOT EDIT.
package menu

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ MenuEntity = &Menu{}

type MenuEntity interface {
	database.Model
	Active() bool
	SetActive(active bool)
	Sorting() uint
	SetSorting(sorting uint)
	SlugID() *string
	SetSlugID(slugID *string)
	ExternalLink() *string
	SetExternalLink(externalLink *string)
	InternalLink() *string
	SetInternalLink(internalLink *string)
	ParentID() *string
	SetParentID(parentID *string)
}

// TableName returns the table name that belongs to the current model.
func (m *Menu) TableName() string {
	return "Menu"
}

// ID returns id.
func (m *Menu) ID() string {
	return m.id
}

// CreatedByID returns createdByID.
func (m *Menu) CreatedByID() string {
	return m.createdByID
}

// SetCreatedByID sets the createdByID.
func (m *Menu) SetCreatedByID(createdByID string) {
	m.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (m *Menu) UpdatedByID() *string {
	return m.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (m *Menu) SetUpdatedByID(updatedByID *string) {
	m.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (m *Menu) CreatedAt() time.Time {
	return m.createdAt
}

// SetCreatedAt sets the createdAt.
func (m *Menu) SetCreatedAt(createdAt time.Time) {
	m.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (m *Menu) UpdatedAt() *time.Time {
	return m.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (m *Menu) SetUpdatedAt(updatedAt *time.Time) {
	m.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (m *Menu) CreatedByFirstName() *string {
	return m.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (m *Menu) SetCreatedByFirstName(createdByFirstName *string) {
	m.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (m *Menu) CreatedBySurname() *string {
	return m.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (m *Menu) SetCreatedBySurname(createdBySurname *string) {
	m.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (m *Menu) UpdatedByFirstName() *string {
	return m.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (m *Menu) SetUpdatedByFirstName(updatedByFirstName *string) {
	m.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (m *Menu) UpdatedBySurname() *string {
	return m.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (m *Menu) SetUpdatedBySurname(updatedBySurname *string) {
	m.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (m *Menu) IsUpdated() bool {
	return m.updatedByID != nil
}

// Active returns active.
func (m *Menu) Active() bool {
	return m.active
}

// SetActive sets the active.
func (m *Menu) SetActive(active bool) {
	m.active = active
}

// Sorting returns sorting.
func (m *Menu) Sorting() uint {
	return m.sorting
}

// SetSorting sets the sorting.
func (m *Menu) SetSorting(sorting uint) {
	m.sorting = sorting
}

// SlugID returns slugID.
func (m *Menu) SlugID() *string {
	return m.slugID
}

// SetSlugID sets the slugID.
func (m *Menu) SetSlugID(slugID *string) {
	m.slugID = slugID
}

// ExternalLink returns externalLink.
func (m *Menu) ExternalLink() *string {
	return m.externalLink
}

// SetExternalLink sets the externalLink.
func (m *Menu) SetExternalLink(externalLink *string) {
	m.externalLink = externalLink
}

// InternalLink returns internalLink.
func (m *Menu) InternalLink() *string {
	return m.internalLink
}

// SetInternalLink sets the internalLink.
func (m *Menu) SetInternalLink(internalLink *string) {
	m.internalLink = internalLink
}

// ParentID returns parentID.
func (m *Menu) ParentID() *string {
	return m.parentID
}

// SetParentID sets the parentID.
func (m *Menu) SetParentID(parentID *string) {
	m.parentID = parentID
}

func newMenu() *Menu {
	return &Menu{}
}

// New returns a new instance of MenuEntity.
func NewMenuEntity() MenuEntity {
	return newMenu()
}
