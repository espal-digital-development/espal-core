package discountcode

import (
	"time"
)

// DiscountCode database object.
// @synthesize
type DiscountCode struct {
	// TODO :: Implement and bind for migration/foreign keys/indexes
	// TODO :: Check Percentage OR Amount is at least set
	// TODO :: Per currency/country/site rules?
	// TODO :: Bind uniquely to certain UserGroups?
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	key                string
	maxUses            *uint
	usesCounter        uint
	availableFrom      *time.Time
	availableUntil     *time.Time
	discountPercentage *float32
	discountAmount     *float32
}
