package property

import (
	"time"
)

// Property database object.
// @synthesize
type Property struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	unitID             *string
	active             bool
	sorting            uint
	key                *string
	_type              uint8 // Can't name it `type` as it's a reserved word
	multiLingual       bool

	// PropertyGroupsProperties []PropertyGroupProperty
	// Options                  []Option
}

// TableAlias returns the unique resolved table alias for use in queries.
func (p *Property) TableAlias() string {
	return "pro"
}
