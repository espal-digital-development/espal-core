package wishlist

import (
	"github.com/espal-digital-development/espal-core/database"
)

// WishlistsStore data store.
type WishlistsStore struct {
	selecterDatabase database.Database
}
