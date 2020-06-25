// Code generated by espal-store-synthesizer. DO NOT EDIT.
package shippingwindow

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ ShippingWindowEntity = &ShippingWindow{}

type ShippingWindowEntity interface {
	database.Model
	UserGroupID() *string
	SetUserGroupID(userGroupID *string)
	StartDate() *time.Time
	SetStartDate(startDate *time.Time)
	EndDate() *time.Time
	SetEndDate(endDate *time.Time)
}

// TableName returns the table name that belongs to the current model.
func (s *ShippingWindow) TableName() string {
	return "ShippingWindow"
}

// ID returns id.
func (s *ShippingWindow) ID() string {
	return s.id
}

// CreatedByID returns createdByID.
func (s *ShippingWindow) CreatedByID() string {
	return s.createdByID
}

// SetCreatedByID sets the createdByID.
func (s *ShippingWindow) SetCreatedByID(createdByID string) {
	s.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (s *ShippingWindow) UpdatedByID() *string {
	return s.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (s *ShippingWindow) SetUpdatedByID(updatedByID *string) {
	s.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (s *ShippingWindow) CreatedAt() time.Time {
	return s.createdAt
}

// SetCreatedAt sets the createdAt.
func (s *ShippingWindow) SetCreatedAt(createdAt time.Time) {
	s.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (s *ShippingWindow) UpdatedAt() *time.Time {
	return s.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (s *ShippingWindow) SetUpdatedAt(updatedAt *time.Time) {
	s.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (s *ShippingWindow) CreatedByFirstName() *string {
	return s.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (s *ShippingWindow) SetCreatedByFirstName(createdByFirstName *string) {
	s.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (s *ShippingWindow) CreatedBySurname() *string {
	return s.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (s *ShippingWindow) SetCreatedBySurname(createdBySurname *string) {
	s.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (s *ShippingWindow) UpdatedByFirstName() *string {
	return s.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (s *ShippingWindow) SetUpdatedByFirstName(updatedByFirstName *string) {
	s.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (s *ShippingWindow) UpdatedBySurname() *string {
	return s.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (s *ShippingWindow) SetUpdatedBySurname(updatedBySurname *string) {
	s.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (s *ShippingWindow) IsUpdated() bool {
	return s.updatedByID != nil
}

// UserGroupID returns userGroupID.
func (s *ShippingWindow) UserGroupID() *string {
	return s.userGroupID
}

// SetUserGroupID sets the userGroupID.
func (s *ShippingWindow) SetUserGroupID(userGroupID *string) {
	s.userGroupID = userGroupID
}

// StartDate returns startDate.
func (s *ShippingWindow) StartDate() *time.Time {
	return s.startDate
}

// SetStartDate sets the startDate.
func (s *ShippingWindow) SetStartDate(startDate *time.Time) {
	s.startDate = startDate
}

// EndDate returns endDate.
func (s *ShippingWindow) EndDate() *time.Time {
	return s.endDate
}

// SetEndDate sets the endDate.
func (s *ShippingWindow) SetEndDate(endDate *time.Time) {
	s.endDate = endDate
}

func newShippingWindow() *ShippingWindow {
	return &ShippingWindow{}
}

// New returns a new instance of ShippingWindowEntity.
func NewShippingWindowEntity() ShippingWindowEntity {
	return newShippingWindow()
}
