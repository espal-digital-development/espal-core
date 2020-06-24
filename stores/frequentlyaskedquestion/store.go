package frequentlyaskedquestion

import (
	"github.com/espal-digital-development/espal-core/database"
)

// FrequentlyAskedQuestionsStore data store.
type FrequentlyAskedQuestionsStore struct {
	selecterDatabase database.Database
}
