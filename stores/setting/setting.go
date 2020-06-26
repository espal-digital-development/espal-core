package setting

import (
	"time"
)

// Setting constants are used for Setting Keys.
// Access order is: User < Domain < Site.
// Should not set DomainID/SiteID/UserID all to NULL.
const (
	// Levels: Domain, Site
	SettingUserSiteAccess       uint16 = 1
	SettingUserSiteAccessGlobal string = "1"
	SettingUserSiteAccessStrict string = "2"
)

// Setting database object.
// @synthesize
type Setting struct {
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	domainID           *string
	siteID             *string
	userID             *string
	key                uint16
	value              string
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *Setting) TableAlias() string {
	return "set"
}
