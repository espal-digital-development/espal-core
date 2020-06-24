// Code generated by espal-store-synthesizer. DO NOT EDIT.
package contact

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ ContactEntity = &Contact{}

type ContactEntity interface {
	database.Model
	UserID() *string
	SetUserID(userID *string)
	ContactID() string
	SetContactID(contactID string)
	Sorting() uint
	SetSorting(sorting uint)
	Comments() *string
	SetComments(comments *string)
	ContactFirstName() *string
	SetContactFirstName(contactFirstName *string)
	ContactSurname() *string
	SetContactSurname(contactSurname *string)
}

// TableAlias returns the unique resolved table alias for use in queries.
func (contact *Contact) TableAlias() string {
	return "ce"
}

// ID returns id.
func (contact *Contact) ID() string {
	return contact.id
}

// CreatedByID returns createdByID.
func (contact *Contact) CreatedByID() string {
	return contact.createdByID
}

// SetCreatedByID sets the createdByID.
func (contact *Contact) SetCreatedByID(createdByID string) {
	contact.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (contact *Contact) UpdatedByID() *string {
	return contact.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (contact *Contact) SetUpdatedByID(updatedByID *string) {
	contact.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (contact *Contact) CreatedAt() time.Time {
	return contact.createdAt
}

// SetCreatedAt sets the createdAt.
func (contact *Contact) SetCreatedAt(createdAt time.Time) {
	contact.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (contact *Contact) UpdatedAt() *time.Time {
	return contact.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (contact *Contact) SetUpdatedAt(updatedAt *time.Time) {
	contact.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (contact *Contact) CreatedByFirstName() *string {
	return contact.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (contact *Contact) SetCreatedByFirstName(createdByFirstName *string) {
	contact.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (contact *Contact) CreatedBySurname() *string {
	return contact.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (contact *Contact) SetCreatedBySurname(createdBySurname *string) {
	contact.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (contact *Contact) UpdatedByFirstName() *string {
	return contact.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (contact *Contact) SetUpdatedByFirstName(updatedByFirstName *string) {
	contact.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (contact *Contact) UpdatedBySurname() *string {
	return contact.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (contact *Contact) SetUpdatedBySurname(updatedBySurname *string) {
	contact.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (contact *Contact) IsUpdated() bool {
	return contact.updatedByID != nil
}

// UserID returns userID.
func (contact *Contact) UserID() *string {
	return contact.userID
}

// SetUserID sets the userID.
func (contact *Contact) SetUserID(userID *string) {
	contact.userID = userID
}

// ContactID returns contactID.
func (contact *Contact) ContactID() string {
	return contact.contactID
}

// SetContactID sets the contactID.
func (contact *Contact) SetContactID(contactID string) {
	contact.contactID = contactID
}

// Sorting returns sorting.
func (contact *Contact) Sorting() uint {
	return contact.sorting
}

// SetSorting sets the sorting.
func (contact *Contact) SetSorting(sorting uint) {
	contact.sorting = sorting
}

// Comments returns comments.
func (contact *Contact) Comments() *string {
	return contact.comments
}

// SetComments sets the comments.
func (contact *Contact) SetComments(comments *string) {
	contact.comments = comments
}

// ContactFirstName returns contactFirstName.
func (contact *Contact) ContactFirstName() *string {
	return contact.contactFirstName
}

// SetContactFirstName sets the contactFirstName.
func (contact *Contact) SetContactFirstName(contactFirstName *string) {
	contact.contactFirstName = contactFirstName
}

// ContactSurname returns contactSurname.
func (contact *Contact) ContactSurname() *string {
	return contact.contactSurname
}

// SetContactSurname sets the contactSurname.
func (contact *Contact) SetContactSurname(contactSurname *string) {
	contact.contactSurname = contactSurname
}

func newContact() *Contact {
	return &Contact{}
}

// New returns a new instance of ContactEntity.
func NewContactEntity() ContactEntity {
	return newContact()
}
