// Code generated by espal-store-synthesizer. DO NOT EDIT.
package newsletter

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ TranslationEntity = &Translation{}

type TranslationEntity interface {
	database.TranslationModel
	NewsletterID() string
	SetNewsletterID(newsletterID string)
}

// ID returns id.
func (translation *Translation) ID() string {
	return translation.id
}

// CreatedByID returns createdByID.
func (translation *Translation) CreatedByID() string {
	return translation.createdByID
}

// SetCreatedByID sets the createdByID.
func (translation *Translation) SetCreatedByID(createdByID string) {
	translation.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (translation *Translation) UpdatedByID() *string {
	return translation.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (translation *Translation) SetUpdatedByID(updatedByID *string) {
	translation.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (translation *Translation) CreatedAt() time.Time {
	return translation.createdAt
}

// SetCreatedAt sets the createdAt.
func (translation *Translation) SetCreatedAt(createdAt time.Time) {
	translation.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (translation *Translation) UpdatedAt() *time.Time {
	return translation.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (translation *Translation) SetUpdatedAt(updatedAt *time.Time) {
	translation.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (translation *Translation) CreatedByFirstName() *string {
	return translation.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (translation *Translation) SetCreatedByFirstName(createdByFirstName *string) {
	translation.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (translation *Translation) CreatedBySurname() *string {
	return translation.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (translation *Translation) SetCreatedBySurname(createdBySurname *string) {
	translation.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (translation *Translation) UpdatedByFirstName() *string {
	return translation.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (translation *Translation) SetUpdatedByFirstName(updatedByFirstName *string) {
	translation.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (translation *Translation) UpdatedBySurname() *string {
	return translation.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (translation *Translation) SetUpdatedBySurname(updatedBySurname *string) {
	translation.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (translation *Translation) IsUpdated() bool {
	return translation.updatedByID != nil
}

// Language returns language.
func (translation *Translation) Language() uint16 {
	return translation.language
}

// SetLanguage sets the language.
func (translation *Translation) SetLanguage(language uint16) {
	translation.language = language
}

// Field returns field.
func (translation *Translation) Field() uint16 {
	return translation.field
}

// SetField sets the field.
func (translation *Translation) SetField(field uint16) {
	translation.field = field
}

// Value returns value.
func (translation *Translation) Value() string {
	return translation.value
}

// SetValue sets the value.
func (translation *Translation) SetValue(value string) {
	translation.value = value
}

// NewsletterID returns newsletterID.
func (translation *Translation) NewsletterID() string {
	return translation.newsletterID
}

// SetNewsletterID sets the newsletterID.
func (translation *Translation) SetNewsletterID(newsletterID string) {
	translation.newsletterID = newsletterID
}

func newTranslation() *Translation {
	return &Translation{}
}

// New returns a new instance of TranslationEntity.
func NewTranslationEntity() TranslationEntity {
	return newTranslation()
}
