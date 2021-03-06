// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ VariantPropertyEntity = &VariantProperty{}

type VariantPropertyEntity interface {
	database.Model
	VariantID() string
	SetVariantID(variantID string)
	PropertyID() string
	SetPropertyID(propertyID string)
	Sorting() uint
	SetSorting(sorting uint)
	Key() *string
	SetKey(key *string)
}

// ID returns id.
func (v *VariantProperty) ID() string {
	return v.id
}

// CreatedByID returns createdByID.
func (v *VariantProperty) CreatedByID() string {
	return v.createdByID
}

// SetCreatedByID sets the createdByID.
func (v *VariantProperty) SetCreatedByID(createdByID string) {
	v.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (v *VariantProperty) UpdatedByID() *string {
	return v.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (v *VariantProperty) SetUpdatedByID(updatedByID *string) {
	v.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (v *VariantProperty) CreatedAt() time.Time {
	return v.createdAt
}

// SetCreatedAt sets the createdAt.
func (v *VariantProperty) SetCreatedAt(createdAt time.Time) {
	v.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (v *VariantProperty) UpdatedAt() *time.Time {
	return v.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (v *VariantProperty) SetUpdatedAt(updatedAt *time.Time) {
	v.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (v *VariantProperty) CreatedByFirstName() *string {
	return v.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (v *VariantProperty) SetCreatedByFirstName(createdByFirstName *string) {
	v.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (v *VariantProperty) CreatedBySurname() *string {
	return v.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (v *VariantProperty) SetCreatedBySurname(createdBySurname *string) {
	v.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (v *VariantProperty) UpdatedByFirstName() *string {
	return v.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (v *VariantProperty) SetUpdatedByFirstName(updatedByFirstName *string) {
	v.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (v *VariantProperty) UpdatedBySurname() *string {
	return v.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (v *VariantProperty) SetUpdatedBySurname(updatedBySurname *string) {
	v.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (v *VariantProperty) IsUpdated() bool {
	return v.updatedByID != nil
}

// VariantID returns variantID.
func (v *VariantProperty) VariantID() string {
	return v.variantID
}

// SetVariantID sets the variantID.
func (v *VariantProperty) SetVariantID(variantID string) {
	v.variantID = variantID
}

// PropertyID returns propertyID.
func (v *VariantProperty) PropertyID() string {
	return v.propertyID
}

// SetPropertyID sets the propertyID.
func (v *VariantProperty) SetPropertyID(propertyID string) {
	v.propertyID = propertyID
}

// Sorting returns sorting.
func (v *VariantProperty) Sorting() uint {
	return v.sorting
}

// SetSorting sets the sorting.
func (v *VariantProperty) SetSorting(sorting uint) {
	v.sorting = sorting
}

// Key returns key.
func (v *VariantProperty) Key() *string {
	return v.key
}

// SetKey sets the key.
func (v *VariantProperty) SetKey(key *string) {
	v.key = key
}

func newVariantProperty() *VariantProperty {
	return &VariantProperty{}
}

// New returns a new instance of VariantPropertyEntity.
func NewVariantPropertyEntity() VariantPropertyEntity {
	return newVariantProperty()
}
