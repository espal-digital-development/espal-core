package page

import (
	"time"
)

// Page database object.
// @synthesize
type Page struct {
	// TODO :: Dynamic data loading into Pages
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
	active             bool
}

// TableAlias returns the unique resolved table alias for use in queries.
func (p *Page) TableAlias() string {
	return "pag"
}
