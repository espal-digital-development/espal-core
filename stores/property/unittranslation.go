package property

import (
	"time"
)

// UnitTranslation database object.
// @synthesize
type UnitTranslation struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	language           uint16
	field              uint16
	unitID             string
	value              string
	display            *string // Optional different display for the translation of the unit
}

// TableName returns the table name that belongs to the current model.
func (unitTranslation *UnitTranslation) TableName() string {
	return "PropertyUnitTranslation"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (unitTranslation *UnitTranslation) TableAlias() string {
	return "prout"
}
