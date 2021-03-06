// Code generated by espal-store-synthesizer. DO NOT EDIT.
package forum

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ PermissionEntity = &Permission{}

type PermissionEntity interface {
	database.Model
	ForumID() *string
	SetForumID(forumID *string)
	UserID() *string
	SetUserID(userID *string)
	Permission() uint8
	SetPermission(permission uint8)
}

// ID returns id.
func (p *Permission) ID() string {
	return p.id
}

// CreatedByID returns createdByID.
func (p *Permission) CreatedByID() string {
	return p.createdByID
}

// SetCreatedByID sets the createdByID.
func (p *Permission) SetCreatedByID(createdByID string) {
	p.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (p *Permission) UpdatedByID() *string {
	return p.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (p *Permission) SetUpdatedByID(updatedByID *string) {
	p.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (p *Permission) CreatedAt() time.Time {
	return p.createdAt
}

// SetCreatedAt sets the createdAt.
func (p *Permission) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (p *Permission) UpdatedAt() *time.Time {
	return p.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (p *Permission) SetUpdatedAt(updatedAt *time.Time) {
	p.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (p *Permission) CreatedByFirstName() *string {
	return p.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (p *Permission) SetCreatedByFirstName(createdByFirstName *string) {
	p.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (p *Permission) CreatedBySurname() *string {
	return p.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (p *Permission) SetCreatedBySurname(createdBySurname *string) {
	p.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (p *Permission) UpdatedByFirstName() *string {
	return p.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (p *Permission) SetUpdatedByFirstName(updatedByFirstName *string) {
	p.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (p *Permission) UpdatedBySurname() *string {
	return p.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (p *Permission) SetUpdatedBySurname(updatedBySurname *string) {
	p.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (p *Permission) IsUpdated() bool {
	return p.updatedByID != nil
}

// ForumID returns forumID.
func (p *Permission) ForumID() *string {
	return p.forumID
}

// SetForumID sets the forumID.
func (p *Permission) SetForumID(forumID *string) {
	p.forumID = forumID
}

// UserID returns userID.
func (p *Permission) UserID() *string {
	return p.userID
}

// SetUserID sets the userID.
func (p *Permission) SetUserID(userID *string) {
	p.userID = userID
}

// Permission returns permission.
func (p *Permission) Permission() uint8 {
	return p.permission
}

// SetPermission sets the permission.
func (p *Permission) SetPermission(permission uint8) {
	p.permission = permission
}

func newPermission() *Permission {
	return &Permission{}
}

// New returns a new instance of PermissionEntity.
func NewPermissionEntity() PermissionEntity {
	return newPermission()
}
