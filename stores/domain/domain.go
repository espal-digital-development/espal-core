package domain

import (
	"strings"
	"time"
)

// nolint:deadcode
type domainMethods interface {
	CurrenciesCount() uint
	HostWithProtocol() string
	HostWithProtocolAndWWW() string
}

// Domain database object.
// @synthesize
type Domain struct {
	// TODO :: Host should not have protocol on it when being saved (http:// or https://)
	id                 string
	createdByID        string
	updatedByID        *string
	siteID             string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	active             bool
	host               string
	language           *uint16
	currencies         string

	localizedName *string // @synthesize-no-db-field
}

// CurrenciesCount returns the amount of currencies that are
// defined in the string field.
func (d *Domain) CurrenciesCount() uint {
	if d.currencies != "" {
		return uint(strings.Count(d.currencies, ",") + 1)
	}
	return 0
}

// HostWithProtocol returns Host with https:// prefixed.
func (d *Domain) HostWithProtocol() string {
	return "https://" + d.host
}

// HostWithProtocolAndWWW returns Host with https://www. prefixed.
func (d *Domain) HostWithProtocolAndWWW() string {
	return "https://www." + d.host
}
