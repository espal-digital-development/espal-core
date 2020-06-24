package invoice

import (
	"time"
)

// Invoice database object.
// @synthesize
type Invoice struct {
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
	saleOrderID               string // TODO :: This probably needs to be SaleOrders, as you can make an Invoice of multiple (partial) SaleOrders
	currency                  uint
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

	// Lines []InvoiceLine
}
