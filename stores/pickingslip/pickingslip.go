package pickingslip

import (
	"time"
)

// PickingSlip database object.
// @synthesize
type PickingSlip struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	comments           *string
}
