package wishlist

import (
	"time"
)

// Wishlist database object.
// @synthesize
type Wishlist struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	domainID           string
	userID             string
	sorting            uint

	// Lines []WishlistLine
}
