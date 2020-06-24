package gift

import (
	"time"
)

// Gift database object.
// @synthesize
type Gift struct {
	// TODO :: Implement (link to ProductVariant and/or ProductBundle?)
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
