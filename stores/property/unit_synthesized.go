// Code generated by espal-store-synthesizer. DO NOT EDIT.
package property

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ UnitEntity = &Unit{}

type UnitEntity interface {
	database.Model
	Display() string
	SetDisplay(display string)
}

// ID returns id.
func (u *Unit) ID() string {
	return u.id
}

// CreatedByID returns createdByID.
func (u *Unit) CreatedByID() string {
	return u.createdByID
}

// SetCreatedByID sets the createdByID.
func (u *Unit) SetCreatedByID(createdByID string) {
	u.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (u *Unit) UpdatedByID() *string {
	return u.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (u *Unit) SetUpdatedByID(updatedByID *string) {
	u.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (u *Unit) CreatedAt() time.Time {
	return u.createdAt
}

// SetCreatedAt sets the createdAt.
func (u *Unit) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (u *Unit) UpdatedAt() *time.Time {
	return u.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (u *Unit) SetUpdatedAt(updatedAt *time.Time) {
	u.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (u *Unit) CreatedByFirstName() *string {
	return u.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (u *Unit) SetCreatedByFirstName(createdByFirstName *string) {
	u.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (u *Unit) CreatedBySurname() *string {
	return u.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (u *Unit) SetCreatedBySurname(createdBySurname *string) {
	u.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (u *Unit) UpdatedByFirstName() *string {
	return u.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (u *Unit) SetUpdatedByFirstName(updatedByFirstName *string) {
	u.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (u *Unit) UpdatedBySurname() *string {
	return u.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (u *Unit) SetUpdatedBySurname(updatedBySurname *string) {
	u.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (u *Unit) IsUpdated() bool {
	return u.updatedByID != nil
}

// Display returns display.
func (u *Unit) Display() string {
	return u.display
}

// SetDisplay sets the display.
func (u *Unit) SetDisplay(display string) {
	u.display = display
}

func newUnit() *Unit {
	return &Unit{}
}

// New returns a new instance of UnitEntity.
func NewUnitEntity() UnitEntity {
	return newUnit()
}
