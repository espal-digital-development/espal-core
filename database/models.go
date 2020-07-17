package database

import (
	"time"
)

// modelBase represents the default requirements for any database model's fields.
type modelBase interface {
	ID() string
	UpdatedByID() *string
	SetUpdatedByID(updatedByID *string)

	CreatedByFirstName() *string
	SetCreatedByFirstName(updatedByFirstName *string)
	CreatedBySurname() *string
	SetCreatedBySurname(createdBySurname *string)

	UpdatedByFirstName() *string
	SetUpdatedByFirstName(updatedByFirstName *string)
	UpdatedBySurname() *string
	SetUpdatedBySurname(updatedBySurname *string)

	CreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	UpdatedAt() *time.Time
	SetUpdatedAt(updatedAt *time.Time)
	IsUpdated() bool
}

// tableMeta represents an object that needs to provide meta data about the database table.
type tableMeta interface {
	TableName() string
	TableAlias() string
}

// modelBaseRequireCreator represents the creator field to be required.
type modelBaseRequireCreator interface {
	CreatedByID() string
	SetCreatedByID(createdByID string)
}

// modelBaseRequireCreator represents the creator field to be optional.
type modelBaseOptionalCreator interface {
	CreatedByID() *string
	SetCreatedByID(createdByID *string)
}

// Model represents the default requirements for any database model's fields where the creator field is required
// (not nullable).
type Model interface {
	tableMeta
	modelBase
	modelBaseRequireCreator
}

// ModelWithOptionalCreator represents the default requirements for any database model's fields where the creator field
// is not required (nullable).
type ModelWithOptionalCreator interface {
	tableMeta
	modelBase
	modelBaseOptionalCreator
}

// TranslationModel represents the default requirements for any database translation model's fields where the creator
// field is required (not nullable).
type TranslationModel interface {
	Model
	Language() uint16
	SetLanguage(language uint16)
	Field() uint16
	SetField(field uint16)
	Value() string
	SetValue(value string)
}
