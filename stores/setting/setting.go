package setting

import (
	"time"
)

// Setting constants are used for Setting Keys.
// Access order is: User < Domain < Site.
// Should not set DomainID/SiteID/UserID all to NULL.
// Keep this list in order as the iota counter is crucial for the reserving of the unique keys.
const (
	// Levels: Domain, Site
	SettingUserSiteAccess uint16 = iota + 1
	SettingTheme
)

const (
	SettingUserSiteAccessGlobal string = "1"
	SettingUserSiteAccessStrict string = "2"
	SettingThemeDefault         string = "default"
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
