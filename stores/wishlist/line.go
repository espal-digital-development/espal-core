package wishlist

import (
	"time"
)

// Line database object.
// @synthesize
type Line struct {
	// TODO :: PreSave only allowed Variant OR Bundled
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	wishlistID         string
	productVariantID   *string
	bundledProductID   *string
	sorting            uint
	quantity           uint
}

// TableName returns the table name that belongs to the current model.
func (line *Line) TableName() string {
	return "WishlistLine"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (line *Line) TableAlias() string {
	return "wl"
}
