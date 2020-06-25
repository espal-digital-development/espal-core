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
func (r *UserRights) GetName(code uint16) (string, error) {
	if r.entries[code] == "" {
		return "", errors.Errorf("userright `%d` doesn't exist", code)
	}
	return r.entries[code], nil
}

// GetCode returns a specific UserRight code based on the requested name.
func (r *UserRights) GetCode(name string) (uint16, error) {
	if r.entriesByName[name] == 0 {
		return 0, errors.Errorf("userright `%s` doesn't exist", name)
	}
	return r.entriesByName[name], nil
}

// AllByCode returns all known UserRights with uint16 index.
func (r *UserRights) AllByCode() map[uint16]string {
	return r.entries
}

// AllByName returns all known UserRights with string index.
func (r *UserRights) AllByName() map[string]uint16 {
	return r.entriesByName
}

// UserRightCodes returns a slice all known UserRights codes in uint16.
func (r *UserRights) UserRightCodes() []uint16 {
	return r.codes
}

// New returns a new instance of r.
func New() *UserRights {
	// Everything under 3e4 is for auto-generation
	// Everything over 3e4 is for manual addition
	r := &UserRights{
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
	for k := range r.entriesSubjectsForCRUD {
		for k2 := range actions {
			r.entriesByName[actions[k2]+r.entriesSubjectsForCRUD[k]] = userRightCode
			r.entries[userRightCode] = actions[k2] + r.entriesSubjectsForCRUD[k]
			userRightCode++
		}
	}

	userRightCodesInts := make([]int, len(r.entries))

	var c int
	for userRight := range r.entries {
		userRightCodesInts[c] = int(userRight)
		c++
	}

	sort.Ints(userRightCodesInts)

	r.codes = make([]uint16, len(userRightCodesInts))

	for k := range userRightCodesInts {
		r.codes[k] = uint16(userRightCodesInts[k])
	}

	return r
}
