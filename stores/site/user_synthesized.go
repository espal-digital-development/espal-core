// Code generated by espal-store-synthesizer. DO NOT EDIT.
package site

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ UserEntity = &User{}

type UserEntity interface {
	database.Model
	SiteID() string
	SetSiteID(siteID string)
	UserID() string
	SetUserID(userID string)
}

// ID returns id.
func (user *User) ID() string {
	return user.id
}

// CreatedByID returns createdByID.
func (user *User) CreatedByID() string {
	return user.createdByID
}

// SetCreatedByID sets the createdByID.
func (user *User) SetCreatedByID(createdByID string) {
	user.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (user *User) UpdatedByID() *string {
	return user.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (user *User) SetUpdatedByID(updatedByID *string) {
	user.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (user *User) CreatedAt() time.Time {
	return user.createdAt
}

// SetCreatedAt sets the createdAt.
func (user *User) SetCreatedAt(createdAt time.Time) {
	user.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (user *User) UpdatedAt() *time.Time {
	return user.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (user *User) SetUpdatedAt(updatedAt *time.Time) {
	user.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (user *User) CreatedByFirstName() *string {
	return user.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (user *User) SetCreatedByFirstName(createdByFirstName *string) {
	user.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (user *User) CreatedBySurname() *string {
	return user.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (user *User) SetCreatedBySurname(createdBySurname *string) {
	user.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (user *User) UpdatedByFirstName() *string {
	return user.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (user *User) SetUpdatedByFirstName(updatedByFirstName *string) {
	user.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (user *User) UpdatedBySurname() *string {
	return user.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (user *User) SetUpdatedBySurname(updatedBySurname *string) {
	user.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (user *User) IsUpdated() bool {
	return user.updatedByID != nil
}

// SiteID returns siteID.
func (user *User) SiteID() string {
	return user.siteID
}

// SetSiteID sets the siteID.
func (user *User) SetSiteID(siteID string) {
	user.siteID = siteID
}

// UserID returns userID.
func (user *User) UserID() string {
	return user.userID
}

// SetUserID sets the userID.
func (user *User) SetUserID(userID string) {
	user.userID = userID
}

func newUser() *User {
	return &User{}
}

// New returns a new instance of UserEntity.
func NewUserEntity() UserEntity {
	return newUser()
}
