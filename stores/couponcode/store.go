package couponcode

import (
	"github.com/espal-digital-development/espal-core/database"
)

// CouponCodesStore data store.
type CouponCodesStore struct {
	selecterDatabase database.Database
}
