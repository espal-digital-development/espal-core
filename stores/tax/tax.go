package tax

import (
	"time"
)

// Tax database object.
// @synthesize
type Tax struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	taxGroupID         string
	country            uint16
	rate               float32
}

// TableAlias returns the unique resolved table alias for use in queries.
func (tax *Tax) TableAlias() string {
	return "ta"
}
