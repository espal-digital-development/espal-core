package discountcode

import (
	"github.com/espal-digital-development/espal-core/database"
)

// DiscountCodesStore data store.
type DiscountCodesStore struct {
	selecterDatabase database.Database
}
