package couponcode

import (
	"time"
)

// CouponCode database object.
// @synthesize
type CouponCode struct {
	// TODO :: Implement and bind for migration/foreign keys/indexes
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	claimedByID        *string // User
	key                string
}
