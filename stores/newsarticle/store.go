package newsarticle

import (
	"github.com/espal-digital-development/espal-core/database"
)

// NewsArticlesStore data store.
type NewsArticlesStore struct {
	selecterDatabase database.Database
}
