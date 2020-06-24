package blogpost

import (
	"github.com/espal-digital-development/espal-core/database"
)

// BlogPostsStore data store.
type BlogPostsStore struct {
	selecterDatabase database.Database
}
