package download

import (
	"time"
)

// Download database object.
// @synthesize
type Download struct {
	// TODO :: More fields?
	// TODO :: Link to unique digital products (e.g. ebook-ed792a08b.pdf).
	// TODO :: Logs in Tiedot of who downloaded what?
	// TODO :: Serve private or public?
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	active             bool
	filePath           string
}

// TableAlias returns the unique resolved table alias for use in queries.
func (download *Download) TableAlias() string {
	return "do"
}
