package saleoffer

import (
	"time"
)

// SaleOffer database object.
// @synthesize
type SaleOffer struct {
	// TODO :: Use M2M Tags to define OfferTypes OrderStatus? (like Presale, Stocksale, Accepted, Canceled, etc.)
	id                               string
	createdByID                      string
	updatedByID                      *string
	createdAt                        time.Time
	updatedAt                        *time.Time
	createdByFirstName               *string
	createdBySurname                 *string
	updatedByFirstName               *string
	updatedBySurname                 *string
	userID                           string
	domainID                         string
	currency                         uint16
	code                             *string
	userInfoBusiness                 bool
	userInfoBusinessCocNumber        *string
	userInfoFirstName                string
	userInfoSurname                  string
	userInfoStreet                   string
	userInfoStreetLine2              *string
	userInfoNumber                   string
	userInfoNumberAddition           *string
	userInfoZipCode                  string
	userInfoCity                     string
	userInfoState                    *uint // TODO :: Store all region-types according to ISO
	userInfoCountry                  *uint16
	userInfoPhoneNumber              *string
	userInfoEmail                    *string
	shippingAddressBusiness          bool
	shippingAddressBusinessCocNumber *string
	shippingAddressFirstName         string
	shippingAddressSurname           string
	shippingAddressStreet            string
	shippingAddressStreetLine2       *string
	shippingAddressNumber            string
	shippingAddressNumberAddition    *string
	shippingAddressZipCode           string
	shippingAddressCity              string
	shippingAddressState             *uint // TODO :: Store all region-types according to ISO
	shippingAddressCountry           *uint16
	shippingAddressPhoneNumber       *string
	shippingAddressEmail             *string
	comments                         *string
	sellingPartyAutograph            *string // bytes for autograph image
	buyingPartyAutograph             *string // bytes for autograph image

	// Lines []Line
}

// TableAlias returns the unique resolved table alias for use in queries.
func (s *SaleOffer) TableAlias() string {
	return "sof"
}
