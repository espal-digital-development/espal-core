package shipment

import (
	"time"
)

// Shipment database object.
// @synthesize
type Shipment struct {
	id                        string
	createdByID               string
	updatedByID               *string
	createdAt                 time.Time
	updatedAt                 *time.Time
	createdByFirstName        *string
	createdBySurname          *string
	updatedByFirstName        *string
	updatedBySurname          *string
	domainID                  string
	userID                    string
	saleOrderID               string
	code                      *string
	userInfoBusiness          bool
	userInfoBusinessCocNumber *string
	userInfoFirstName         string
	userInfoSurname           string
	userInfoStreet            string
	userInfoStreetLine2       *string
	userInfoNumber            string
	userInfoNumberAddition    *string
	userInfoZipCode           string
	userInfoCity              string
	userInfoState             *uint // TODO :: Store all region-types according to ISO
	userInfoCountry           *uint16
	userInfoPhoneNumber       *string
	userInfoEmail             *string
	comments                  *string
	sellingPartyAutograph     *string // bytes for autograph image
	buyingPartyAutograph      *string // bytes for autograph image

	// Lines []Line
}

// TableAlias returns the unique resolved table alias for use in queries.
func (shipment *Shipment) TableAlias() string {
	return "sh"
}
