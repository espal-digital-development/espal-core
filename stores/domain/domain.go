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
func (domain *Domain) CurrenciesCount() uint {
	if domain.currencies != "" {
		return uint(strings.Count(domain.currencies, ",") + 1)
	}
	return 0
}

// HostWithProtocol returns Host with https:// prefixed.
func (domain *Domain) HostWithProtocol() string {
	return "https://" + domain.host
}

// HostWithProtocolAndWWW returns Host with https://www. prefixed.
func (domain *Domain) HostWithProtocolAndWWW() string {
	return "https://www." + domain.host
}
