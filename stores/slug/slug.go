package slug

import (
	"time"
)

// TODO :: Auto-rules for new entities?
// page/:id > nl > algemene-voorwaarden > en > terms-of-agreement
// office/:id > nl > kantoor/groningen-{id} > en > office/groningen-{id}
// TODO :: Sluggify method
// TODO :: Need some constant integrity check to see if the internal routes exist?

// Slug database object.
// @synthesize
type Slug struct {
	// TODO :: Disallow registered routes on Save and Update,
	// also the config.AdminURL(), and also when the config on that gets changed.
	// Maybe add Expiration Time or Date and the action that will trigger then?
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	domainID           string
	language           uint16
	path               string
	rerouteTo          string
	invalidWithStatus  *uint16
	invalidMessage     *string
	// Can be an internal or external fully qualified path
	redirectToRawPath  *string
	redirectStatusCode *uint16
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *Slug) TableAlias() string {
	return "slu"
}
