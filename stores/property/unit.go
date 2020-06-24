package property

import (
	"time"
)

// Unit database object.
// @synthesize
type Unit struct {
	// TODO :: Conversion service of display units (Kg to Pounds, € to $)
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	display            string // like x (qty), MHz, PWu, Volts, Kg, Pounds, MB, GB, µs, €, $
}

// TableName returns the table name that belongs to the current model.
func (unit *Unit) TableName() string {
	return "PropertyUnit"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (unit *Unit) TableAlias() string {
	return "prou"
}
