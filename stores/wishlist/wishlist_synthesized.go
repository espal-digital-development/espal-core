// Code generated by espal-store-synthesizer. DO NOT EDIT.
package wishlist

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ WishlistEntity = &Wishlist{}

type WishlistEntity interface {
	database.Model
	DomainID() string
	SetDomainID(domainID string)
	UserID() string
	SetUserID(userID string)
	Sorting() uint
	SetSorting(sorting uint)
}

// TableName returns the table name that belongs to the current model.
func (w *Wishlist) TableName() string {
	return "Wishlist"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (w *Wishlist) TableAlias() string {
	return "we"
}

// ID returns id.
func (w *Wishlist) ID() string {
	return w.id
}

// CreatedByID returns createdByID.
func (w *Wishlist) CreatedByID() string {
	return w.createdByID
}

// SetCreatedByID sets the createdByID.
func (w *Wishlist) SetCreatedByID(createdByID string) {
	w.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (w *Wishlist) UpdatedByID() *string {
	return w.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (w *Wishlist) SetUpdatedByID(updatedByID *string) {
	w.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (w *Wishlist) CreatedAt() time.Time {
	return w.createdAt
}

// SetCreatedAt sets the createdAt.
func (w *Wishlist) SetCreatedAt(createdAt time.Time) {
	w.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (w *Wishlist) UpdatedAt() *time.Time {
	return w.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (w *Wishlist) SetUpdatedAt(updatedAt *time.Time) {
	w.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (w *Wishlist) CreatedByFirstName() *string {
	return w.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (w *Wishlist) SetCreatedByFirstName(createdByFirstName *string) {
	w.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (w *Wishlist) CreatedBySurname() *string {
	return w.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (w *Wishlist) SetCreatedBySurname(createdBySurname *string) {
	w.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (w *Wishlist) UpdatedByFirstName() *string {
	return w.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (w *Wishlist) SetUpdatedByFirstName(updatedByFirstName *string) {
	w.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (w *Wishlist) UpdatedBySurname() *string {
	return w.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (w *Wishlist) SetUpdatedBySurname(updatedBySurname *string) {
	w.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (w *Wishlist) IsUpdated() bool {
	return w.updatedByID != nil
}

// DomainID returns domainID.
func (w *Wishlist) DomainID() string {
	return w.domainID
}

// SetDomainID sets the domainID.
func (w *Wishlist) SetDomainID(domainID string) {
	w.domainID = domainID
}

// UserID returns userID.
func (w *Wishlist) UserID() string {
	return w.userID
}

// SetUserID sets the userID.
func (w *Wishlist) SetUserID(userID string) {
	w.userID = userID
}

// Sorting returns sorting.
func (w *Wishlist) Sorting() uint {
	return w.sorting
}

// SetSorting sets the sorting.
func (w *Wishlist) SetSorting(sorting uint) {
	w.sorting = sorting
}

func newWishlist() *Wishlist {
	return &Wishlist{}
}

// New returns a new instance of WishlistEntity.
func NewWishlistEntity() WishlistEntity {
	return newWishlist()
}
