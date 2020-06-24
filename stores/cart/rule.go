package cart

import (
	"time"
)

// Rule database object.
// These are the "rules" for a Cart, not the Lines.
// @synthesize
type Rule struct {
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
}

// TableName returns the table name that belongs to the current model.
func (rule *Rule) TableName() string {
	return "CartRule"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (rule *Rule) TableAlias() string {
	return "cr"
}
