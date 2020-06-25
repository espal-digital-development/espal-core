// Code generated by espal-store-synthesizer. DO NOT EDIT.
package frequentlyaskedquestion

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ TranslationEntity = &Translation{}

type TranslationEntity interface {
	database.TranslationModel
	FrequentlyAskedQuestionID() string
	SetFrequentlyAskedQuestionID(frequentlyAskedQuestionID string)
}

// ID returns id.
func (t *Translation) ID() string {
	return t.id
}

// CreatedByID returns createdByID.
func (t *Translation) CreatedByID() string {
	return t.createdByID
}

// SetCreatedByID sets the createdByID.
func (t *Translation) SetCreatedByID(createdByID string) {
	t.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (t *Translation) UpdatedByID() *string {
	return t.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (t *Translation) SetUpdatedByID(updatedByID *string) {
	t.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (t *Translation) CreatedAt() time.Time {
	return t.createdAt
}

// SetCreatedAt sets the createdAt.
func (t *Translation) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (t *Translation) UpdatedAt() *time.Time {
	return t.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (t *Translation) SetUpdatedAt(updatedAt *time.Time) {
	t.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (t *Translation) CreatedByFirstName() *string {
	return t.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (t *Translation) SetCreatedByFirstName(createdByFirstName *string) {
	t.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (t *Translation) CreatedBySurname() *string {
	return t.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (t *Translation) SetCreatedBySurname(createdBySurname *string) {
	t.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (t *Translation) UpdatedByFirstName() *string {
	return t.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (t *Translation) SetUpdatedByFirstName(updatedByFirstName *string) {
	t.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (t *Translation) UpdatedBySurname() *string {
	return t.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (t *Translation) SetUpdatedBySurname(updatedBySurname *string) {
	t.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (t *Translation) IsUpdated() bool {
	return t.updatedByID != nil
}

// Language returns language.
func (t *Translation) Language() uint16 {
	return t.language
}

// SetLanguage sets the language.
func (t *Translation) SetLanguage(language uint16) {
	t.language = language
}

// Field returns field.
func (t *Translation) Field() uint16 {
	return t.field
}

// SetField sets the field.
func (t *Translation) SetField(field uint16) {
	t.field = field
}

// Value returns value.
func (t *Translation) Value() string {
	return t.value
}

// SetValue sets the value.
func (t *Translation) SetValue(value string) {
	t.value = value
}

// FrequentlyAskedQuestionID returns frequentlyAskedQuestionID.
func (t *Translation) FrequentlyAskedQuestionID() string {
	return t.frequentlyAskedQuestionID
}

// SetFrequentlyAskedQuestionID sets the frequentlyAskedQuestionID.
func (t *Translation) SetFrequentlyAskedQuestionID(frequentlyAskedQuestionID string) {
	t.frequentlyAskedQuestionID = frequentlyAskedQuestionID
}

func newTranslation() *Translation {
	return &Translation{}
}

// New returns a new instance of TranslationEntity.
func NewTranslationEntity() TranslationEntity {
	return newTranslation()
}
