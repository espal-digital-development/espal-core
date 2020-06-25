package menu

import (
	"time"
)

// Menu database object.
// @synthesize
type Menu struct {
	// TODO :: On PreSave/PreUpdate check if combinations are valid.
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
	sorting            uint
	slugID             *string
	externalLink       *string
	internalLink       *string // TODO :: Use a light-weight shortcut for internal routes?
	parentID           *string // Menu

	// Children []Menu
}

// TableAlias returns the unique resolved table alias for use in queries.
func (m *Menu) TableAlias() string {
	return "men"
}
