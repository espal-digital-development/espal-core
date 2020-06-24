package userrights

import (
	"sort"

	"github.com/juju/errors"
)

var _ Repository = &UserRights{}

// Repository represents a UserRight repository.
type Repository interface {
	GetName(code uint16) (string, error)
	GetCode(name string) (uint16, error)
	AllByCode() map[uint16]string
	AllByName() map[string]uint16
	UserRightCodes() []uint16
}

// UserRights contains a full UserRight repository.
type UserRights struct {
	codes                  []uint16
	entries                map[uint16]string
	entriesByName          map[string]uint16
	entriesSubjectsForCRUD []string
}

// GetName returns a specific UserRight name based on the requested code.
func (userRights *UserRights) GetName(code uint16) (string, error) {
	if userRights.entries[code] == "" {
		return "", errors.Errorf("userright `%d` doesn't exist", code)
	}
	return userRights.entries[code], nil
}

// GetCode returns a specific UserRight code based on the requested name.
func (userRights *UserRights) GetCode(name string) (uint16, error) {
	if userRights.entriesByName[name] == 0 {
		return 0, errors.Errorf("userright `%s` doesn't exist", name)
	}
	return userRights.entriesByName[name], nil
}

// AllByCode returns all known UserRights with uint16 index.
func (userRights *UserRights) AllByCode() map[uint16]string {
	return userRights.entries
}

// AllByName returns all known UserRights with string index.
func (userRights *UserRights) AllByName() map[string]uint16 {
	return userRights.entriesByName
}

// UserRightCodes returns a slice all known UserRights codes in uint16.
func (userRights *UserRights) UserRightCodes() []uint16 {
	return userRights.codes
}

// New returns a new instance of UserRights.
func New() *UserRights {
	// Everything under 3e4 is for auto-generation
	// Everything over 3e4 is for manual addition
	userRights := &UserRights{
		entries: map[uint16]string{
			30000: "AccessAuth",
			30001: "AccessAdminSection",
		},
		entriesByName: map[string]uint16{
			"AccessAuth":         30000,
			"AccessAdminSection": 30001,
		},
		entriesSubjectsForCRUD: []string{
			"ProductModel", "ProductVariant", "BundledProduct", "Subscription",
			"PropertyGroup", "Property", "ProductReview",
			"Filter", "Cart", "Wishlist", "DiscountCode", "CouponCode", "Credit",
			"SaleOrder", "Invoice", "PurchaseOrder", "PaymentTransaction",
			"PaymentAccount", "PaymentProvider", "PriceGroup", "PriceList",
			"ShipmentCost", "Tax", "Lead", "Opportunity", "Account", "Contact",
			"Report", "Group", "Person", "Offer", "Prospect", "Campaign",
			"Project", "Shipment", "ReturnOrder", "Receiving", "PickingSlip",
			"ShippingWindow", "DeliveryMethod", "Stock", "StockLocation",
			"Supplier", "Menu", "Page", "Block", "Poll", "NewsArticle",
			"BlogPost", "Forum", "Office", "Reseller", "FrequentlyAskedQuestion",
			"EmailTemplate", "Newsletter", "Gift", "GiftWrapping", "Media",
			"Download", "Phase", "Prototype", "Sizescreen", "BillOfMaterial",
			"ImportProfile", "ExportProfile", "SwimmingLane", "DataMapping",
			"Webservice", "ReportAndLog", "AccessLog", "ErrorLog", "SearchHistory",
			"EmailLog", "User", "UserAddress", "UserContact", "UserGroup",
			"UserTask", "Shop", "Slug", "Tag", "Setting", "TechnicalStatistic",
			"Domain", "Site",
		},
	}

	var userRightCode uint16 = 1
	actions := []string{"Access", "Read", "Create", "Update", "Delete"}
	for k := range userRights.entriesSubjectsForCRUD {
		for k2 := range actions {
			userRights.entriesByName[actions[k2]+userRights.entriesSubjectsForCRUD[k]] = userRightCode
			userRights.entries[userRightCode] = actions[k2] + userRights.entriesSubjectsForCRUD[k]
			userRightCode++
		}
	}

	userRightCodesInts := make([]int, len(userRights.entries))

	var c int
	for userRight := range userRights.entries {
		userRightCodesInts[c] = int(userRight)
		c++
	}

	sort.Ints(userRightCodesInts)

	userRights.codes = make([]uint16, len(userRightCodesInts))

	for k := range userRightCodesInts {
		userRights.codes[k] = uint16(userRightCodesInts[k])
	}

	return userRights
}
