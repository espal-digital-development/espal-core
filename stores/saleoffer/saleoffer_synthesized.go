// Code generated by espal-store-synthesizer. DO NOT EDIT.
package saleoffer

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ SaleOfferEntity = &SaleOffer{}

type SaleOfferEntity interface {
	database.Model
	UserID() string
	SetUserID(userID string)
	DomainID() string
	SetDomainID(domainID string)
	Currency() uint16
	SetCurrency(currency uint16)
	Code() *string
	SetCode(code *string)
	UserInfoBusiness() bool
	SetUserInfoBusiness(userInfoBusiness bool)
	UserInfoBusinessCocNumber() *string
	SetUserInfoBusinessCocNumber(userInfoBusinessCocNumber *string)
	UserInfoFirstName() string
	SetUserInfoFirstName(userInfoFirstName string)
	UserInfoSurname() string
	SetUserInfoSurname(userInfoSurname string)
	UserInfoStreet() string
	SetUserInfoStreet(userInfoStreet string)
	UserInfoStreetLine2() *string
	SetUserInfoStreetLine2(userInfoStreetLine2 *string)
	UserInfoNumber() string
	SetUserInfoNumber(userInfoNumber string)
	UserInfoNumberAddition() *string
	SetUserInfoNumberAddition(userInfoNumberAddition *string)
	UserInfoZipCode() string
	SetUserInfoZipCode(userInfoZipCode string)
	UserInfoCity() string
	SetUserInfoCity(userInfoCity string)
	UserInfoState() *uint
	SetUserInfoState(userInfoState *uint)
	UserInfoCountry() *uint16
	SetUserInfoCountry(userInfoCountry *uint16)
	UserInfoPhoneNumber() *string
	SetUserInfoPhoneNumber(userInfoPhoneNumber *string)
	UserInfoEmail() *string
	SetUserInfoEmail(userInfoEmail *string)
	ShippingAddressBusiness() bool
	SetShippingAddressBusiness(shippingAddressBusiness bool)
	ShippingAddressBusinessCocNumber() *string
	SetShippingAddressBusinessCocNumber(shippingAddressBusinessCocNumber *string)
	ShippingAddressFirstName() string
	SetShippingAddressFirstName(shippingAddressFirstName string)
	ShippingAddressSurname() string
	SetShippingAddressSurname(shippingAddressSurname string)
	ShippingAddressStreet() string
	SetShippingAddressStreet(shippingAddressStreet string)
	ShippingAddressStreetLine2() *string
	SetShippingAddressStreetLine2(shippingAddressStreetLine2 *string)
	ShippingAddressNumber() string
	SetShippingAddressNumber(shippingAddressNumber string)
	ShippingAddressNumberAddition() *string
	SetShippingAddressNumberAddition(shippingAddressNumberAddition *string)
	ShippingAddressZipCode() string
	SetShippingAddressZipCode(shippingAddressZipCode string)
	ShippingAddressCity() string
	SetShippingAddressCity(shippingAddressCity string)
	ShippingAddressState() *uint
	SetShippingAddressState(shippingAddressState *uint)
	ShippingAddressCountry() *uint16
	SetShippingAddressCountry(shippingAddressCountry *uint16)
	ShippingAddressPhoneNumber() *string
	SetShippingAddressPhoneNumber(shippingAddressPhoneNumber *string)
	ShippingAddressEmail() *string
	SetShippingAddressEmail(shippingAddressEmail *string)
	Comments() *string
	SetComments(comments *string)
	SellingPartyAutograph() *string
	SetSellingPartyAutograph(sellingPartyAutograph *string)
	BuyingPartyAutograph() *string
	SetBuyingPartyAutograph(buyingPartyAutograph *string)
}

// TableName returns the table name that belongs to the current model.
func (saleOffer *SaleOffer) TableName() string {
	return "SaleOffer"
}

// ID returns id.
func (saleOffer *SaleOffer) ID() string {
	return saleOffer.id
}

// CreatedByID returns createdByID.
func (saleOffer *SaleOffer) CreatedByID() string {
	return saleOffer.createdByID
}

// SetCreatedByID sets the createdByID.
func (saleOffer *SaleOffer) SetCreatedByID(createdByID string) {
	saleOffer.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (saleOffer *SaleOffer) UpdatedByID() *string {
	return saleOffer.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (saleOffer *SaleOffer) SetUpdatedByID(updatedByID *string) {
	saleOffer.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (saleOffer *SaleOffer) CreatedAt() time.Time {
	return saleOffer.createdAt
}

// SetCreatedAt sets the createdAt.
func (saleOffer *SaleOffer) SetCreatedAt(createdAt time.Time) {
	saleOffer.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (saleOffer *SaleOffer) UpdatedAt() *time.Time {
	return saleOffer.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (saleOffer *SaleOffer) SetUpdatedAt(updatedAt *time.Time) {
	saleOffer.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (saleOffer *SaleOffer) CreatedByFirstName() *string {
	return saleOffer.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (saleOffer *SaleOffer) SetCreatedByFirstName(createdByFirstName *string) {
	saleOffer.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (saleOffer *SaleOffer) CreatedBySurname() *string {
	return saleOffer.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (saleOffer *SaleOffer) SetCreatedBySurname(createdBySurname *string) {
	saleOffer.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (saleOffer *SaleOffer) UpdatedByFirstName() *string {
	return saleOffer.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (saleOffer *SaleOffer) SetUpdatedByFirstName(updatedByFirstName *string) {
	saleOffer.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (saleOffer *SaleOffer) UpdatedBySurname() *string {
	return saleOffer.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (saleOffer *SaleOffer) SetUpdatedBySurname(updatedBySurname *string) {
	saleOffer.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (saleOffer *SaleOffer) IsUpdated() bool {
	return saleOffer.updatedByID != nil
}

// UserID returns userID.
func (saleOffer *SaleOffer) UserID() string {
	return saleOffer.userID
}

// SetUserID sets the userID.
func (saleOffer *SaleOffer) SetUserID(userID string) {
	saleOffer.userID = userID
}

// DomainID returns domainID.
func (saleOffer *SaleOffer) DomainID() string {
	return saleOffer.domainID
}

// SetDomainID sets the domainID.
func (saleOffer *SaleOffer) SetDomainID(domainID string) {
	saleOffer.domainID = domainID
}

// Currency returns currency.
func (saleOffer *SaleOffer) Currency() uint16 {
	return saleOffer.currency
}

// SetCurrency sets the currency.
func (saleOffer *SaleOffer) SetCurrency(currency uint16) {
	saleOffer.currency = currency
}

// Code returns code.
func (saleOffer *SaleOffer) Code() *string {
	return saleOffer.code
}

// SetCode sets the code.
func (saleOffer *SaleOffer) SetCode(code *string) {
	saleOffer.code = code
}

// UserInfoBusiness returns userInfoBusiness.
func (saleOffer *SaleOffer) UserInfoBusiness() bool {
	return saleOffer.userInfoBusiness
}

// SetUserInfoBusiness sets the userInfoBusiness.
func (saleOffer *SaleOffer) SetUserInfoBusiness(userInfoBusiness bool) {
	saleOffer.userInfoBusiness = userInfoBusiness
}

// UserInfoBusinessCocNumber returns userInfoBusinessCocNumber.
func (saleOffer *SaleOffer) UserInfoBusinessCocNumber() *string {
	return saleOffer.userInfoBusinessCocNumber
}

// SetUserInfoBusinessCocNumber sets the userInfoBusinessCocNumber.
func (saleOffer *SaleOffer) SetUserInfoBusinessCocNumber(userInfoBusinessCocNumber *string) {
	saleOffer.userInfoBusinessCocNumber = userInfoBusinessCocNumber
}

// UserInfoFirstName returns userInfoFirstName.
func (saleOffer *SaleOffer) UserInfoFirstName() string {
	return saleOffer.userInfoFirstName
}

// SetUserInfoFirstName sets the userInfoFirstName.
func (saleOffer *SaleOffer) SetUserInfoFirstName(userInfoFirstName string) {
	saleOffer.userInfoFirstName = userInfoFirstName
}

// UserInfoSurname returns userInfoSurname.
func (saleOffer *SaleOffer) UserInfoSurname() string {
	return saleOffer.userInfoSurname
}

// SetUserInfoSurname sets the userInfoSurname.
func (saleOffer *SaleOffer) SetUserInfoSurname(userInfoSurname string) {
	saleOffer.userInfoSurname = userInfoSurname
}

// UserInfoStreet returns userInfoStreet.
func (saleOffer *SaleOffer) UserInfoStreet() string {
	return saleOffer.userInfoStreet
}

// SetUserInfoStreet sets the userInfoStreet.
func (saleOffer *SaleOffer) SetUserInfoStreet(userInfoStreet string) {
	saleOffer.userInfoStreet = userInfoStreet
}

// UserInfoStreetLine2 returns userInfoStreetLine2.
func (saleOffer *SaleOffer) UserInfoStreetLine2() *string {
	return saleOffer.userInfoStreetLine2
}

// SetUserInfoStreetLine2 sets the userInfoStreetLine2.
func (saleOffer *SaleOffer) SetUserInfoStreetLine2(userInfoStreetLine2 *string) {
	saleOffer.userInfoStreetLine2 = userInfoStreetLine2
}

// UserInfoNumber returns userInfoNumber.
func (saleOffer *SaleOffer) UserInfoNumber() string {
	return saleOffer.userInfoNumber
}

// SetUserInfoNumber sets the userInfoNumber.
func (saleOffer *SaleOffer) SetUserInfoNumber(userInfoNumber string) {
	saleOffer.userInfoNumber = userInfoNumber
}

// UserInfoNumberAddition returns userInfoNumberAddition.
func (saleOffer *SaleOffer) UserInfoNumberAddition() *string {
	return saleOffer.userInfoNumberAddition
}

// SetUserInfoNumberAddition sets the userInfoNumberAddition.
func (saleOffer *SaleOffer) SetUserInfoNumberAddition(userInfoNumberAddition *string) {
	saleOffer.userInfoNumberAddition = userInfoNumberAddition
}

// UserInfoZipCode returns userInfoZipCode.
func (saleOffer *SaleOffer) UserInfoZipCode() string {
	return saleOffer.userInfoZipCode
}

// SetUserInfoZipCode sets the userInfoZipCode.
func (saleOffer *SaleOffer) SetUserInfoZipCode(userInfoZipCode string) {
	saleOffer.userInfoZipCode = userInfoZipCode
}

// UserInfoCity returns userInfoCity.
func (saleOffer *SaleOffer) UserInfoCity() string {
	return saleOffer.userInfoCity
}

// SetUserInfoCity sets the userInfoCity.
func (saleOffer *SaleOffer) SetUserInfoCity(userInfoCity string) {
	saleOffer.userInfoCity = userInfoCity
}

// UserInfoState returns userInfoState.
func (saleOffer *SaleOffer) UserInfoState() *uint {
	return saleOffer.userInfoState
}

// SetUserInfoState sets the userInfoState.
func (saleOffer *SaleOffer) SetUserInfoState(userInfoState *uint) {
	saleOffer.userInfoState = userInfoState
}

// UserInfoCountry returns userInfoCountry.
func (saleOffer *SaleOffer) UserInfoCountry() *uint16 {
	return saleOffer.userInfoCountry
}

// SetUserInfoCountry sets the userInfoCountry.
func (saleOffer *SaleOffer) SetUserInfoCountry(userInfoCountry *uint16) {
	saleOffer.userInfoCountry = userInfoCountry
}

// UserInfoPhoneNumber returns userInfoPhoneNumber.
func (saleOffer *SaleOffer) UserInfoPhoneNumber() *string {
	return saleOffer.userInfoPhoneNumber
}

// SetUserInfoPhoneNumber sets the userInfoPhoneNumber.
func (saleOffer *SaleOffer) SetUserInfoPhoneNumber(userInfoPhoneNumber *string) {
	saleOffer.userInfoPhoneNumber = userInfoPhoneNumber
}

// UserInfoEmail returns userInfoEmail.
func (saleOffer *SaleOffer) UserInfoEmail() *string {
	return saleOffer.userInfoEmail
}

// SetUserInfoEmail sets the userInfoEmail.
func (saleOffer *SaleOffer) SetUserInfoEmail(userInfoEmail *string) {
	saleOffer.userInfoEmail = userInfoEmail
}

// ShippingAddressBusiness returns shippingAddressBusiness.
func (saleOffer *SaleOffer) ShippingAddressBusiness() bool {
	return saleOffer.shippingAddressBusiness
}

// SetShippingAddressBusiness sets the shippingAddressBusiness.
func (saleOffer *SaleOffer) SetShippingAddressBusiness(shippingAddressBusiness bool) {
	saleOffer.shippingAddressBusiness = shippingAddressBusiness
}

// ShippingAddressBusinessCocNumber returns shippingAddressBusinessCocNumber.
func (saleOffer *SaleOffer) ShippingAddressBusinessCocNumber() *string {
	return saleOffer.shippingAddressBusinessCocNumber
}

// SetShippingAddressBusinessCocNumber sets the shippingAddressBusinessCocNumber.
func (saleOffer *SaleOffer) SetShippingAddressBusinessCocNumber(shippingAddressBusinessCocNumber *string) {
	saleOffer.shippingAddressBusinessCocNumber = shippingAddressBusinessCocNumber
}

// ShippingAddressFirstName returns shippingAddressFirstName.
func (saleOffer *SaleOffer) ShippingAddressFirstName() string {
	return saleOffer.shippingAddressFirstName
}

// SetShippingAddressFirstName sets the shippingAddressFirstName.
func (saleOffer *SaleOffer) SetShippingAddressFirstName(shippingAddressFirstName string) {
	saleOffer.shippingAddressFirstName = shippingAddressFirstName
}

// ShippingAddressSurname returns shippingAddressSurname.
func (saleOffer *SaleOffer) ShippingAddressSurname() string {
	return saleOffer.shippingAddressSurname
}

// SetShippingAddressSurname sets the shippingAddressSurname.
func (saleOffer *SaleOffer) SetShippingAddressSurname(shippingAddressSurname string) {
	saleOffer.shippingAddressSurname = shippingAddressSurname
}

// ShippingAddressStreet returns shippingAddressStreet.
func (saleOffer *SaleOffer) ShippingAddressStreet() string {
	return saleOffer.shippingAddressStreet
}

// SetShippingAddressStreet sets the shippingAddressStreet.
func (saleOffer *SaleOffer) SetShippingAddressStreet(shippingAddressStreet string) {
	saleOffer.shippingAddressStreet = shippingAddressStreet
}

// ShippingAddressStreetLine2 returns shippingAddressStreetLine2.
func (saleOffer *SaleOffer) ShippingAddressStreetLine2() *string {
	return saleOffer.shippingAddressStreetLine2
}

// SetShippingAddressStreetLine2 sets the shippingAddressStreetLine2.
func (saleOffer *SaleOffer) SetShippingAddressStreetLine2(shippingAddressStreetLine2 *string) {
	saleOffer.shippingAddressStreetLine2 = shippingAddressStreetLine2
}

// ShippingAddressNumber returns shippingAddressNumber.
func (saleOffer *SaleOffer) ShippingAddressNumber() string {
	return saleOffer.shippingAddressNumber
}

// SetShippingAddressNumber sets the shippingAddressNumber.
func (saleOffer *SaleOffer) SetShippingAddressNumber(shippingAddressNumber string) {
	saleOffer.shippingAddressNumber = shippingAddressNumber
}

// ShippingAddressNumberAddition returns shippingAddressNumberAddition.
func (saleOffer *SaleOffer) ShippingAddressNumberAddition() *string {
	return saleOffer.shippingAddressNumberAddition
}

// SetShippingAddressNumberAddition sets the shippingAddressNumberAddition.
func (saleOffer *SaleOffer) SetShippingAddressNumberAddition(shippingAddressNumberAddition *string) {
	saleOffer.shippingAddressNumberAddition = shippingAddressNumberAddition
}

// ShippingAddressZipCode returns shippingAddressZipCode.
func (saleOffer *SaleOffer) ShippingAddressZipCode() string {
	return saleOffer.shippingAddressZipCode
}

// SetShippingAddressZipCode sets the shippingAddressZipCode.
func (saleOffer *SaleOffer) SetShippingAddressZipCode(shippingAddressZipCode string) {
	saleOffer.shippingAddressZipCode = shippingAddressZipCode
}

// ShippingAddressCity returns shippingAddressCity.
func (saleOffer *SaleOffer) ShippingAddressCity() string {
	return saleOffer.shippingAddressCity
}

// SetShippingAddressCity sets the shippingAddressCity.
func (saleOffer *SaleOffer) SetShippingAddressCity(shippingAddressCity string) {
	saleOffer.shippingAddressCity = shippingAddressCity
}

// ShippingAddressState returns shippingAddressState.
func (saleOffer *SaleOffer) ShippingAddressState() *uint {
	return saleOffer.shippingAddressState
}

// SetShippingAddressState sets the shippingAddressState.
func (saleOffer *SaleOffer) SetShippingAddressState(shippingAddressState *uint) {
	saleOffer.shippingAddressState = shippingAddressState
}

// ShippingAddressCountry returns shippingAddressCountry.
func (saleOffer *SaleOffer) ShippingAddressCountry() *uint16 {
	return saleOffer.shippingAddressCountry
}

// SetShippingAddressCountry sets the shippingAddressCountry.
func (saleOffer *SaleOffer) SetShippingAddressCountry(shippingAddressCountry *uint16) {
	saleOffer.shippingAddressCountry = shippingAddressCountry
}

// ShippingAddressPhoneNumber returns shippingAddressPhoneNumber.
func (saleOffer *SaleOffer) ShippingAddressPhoneNumber() *string {
	return saleOffer.shippingAddressPhoneNumber
}

// SetShippingAddressPhoneNumber sets the shippingAddressPhoneNumber.
func (saleOffer *SaleOffer) SetShippingAddressPhoneNumber(shippingAddressPhoneNumber *string) {
	saleOffer.shippingAddressPhoneNumber = shippingAddressPhoneNumber
}

// ShippingAddressEmail returns shippingAddressEmail.
func (saleOffer *SaleOffer) ShippingAddressEmail() *string {
	return saleOffer.shippingAddressEmail
}

// SetShippingAddressEmail sets the shippingAddressEmail.
func (saleOffer *SaleOffer) SetShippingAddressEmail(shippingAddressEmail *string) {
	saleOffer.shippingAddressEmail = shippingAddressEmail
}

// Comments returns comments.
func (saleOffer *SaleOffer) Comments() *string {
	return saleOffer.comments
}

// SetComments sets the comments.
func (saleOffer *SaleOffer) SetComments(comments *string) {
	saleOffer.comments = comments
}

// SellingPartyAutograph returns sellingPartyAutograph.
func (saleOffer *SaleOffer) SellingPartyAutograph() *string {
	return saleOffer.sellingPartyAutograph
}

// SetSellingPartyAutograph sets the sellingPartyAutograph.
func (saleOffer *SaleOffer) SetSellingPartyAutograph(sellingPartyAutograph *string) {
	saleOffer.sellingPartyAutograph = sellingPartyAutograph
}

// BuyingPartyAutograph returns buyingPartyAutograph.
func (saleOffer *SaleOffer) BuyingPartyAutograph() *string {
	return saleOffer.buyingPartyAutograph
}

// SetBuyingPartyAutograph sets the buyingPartyAutograph.
func (saleOffer *SaleOffer) SetBuyingPartyAutograph(buyingPartyAutograph *string) {
	saleOffer.buyingPartyAutograph = buyingPartyAutograph
}

func newSaleOffer() *SaleOffer {
	return &SaleOffer{}
}

// New returns a new instance of SaleOfferEntity.
func NewSaleOfferEntity() SaleOfferEntity {
	return newSaleOffer()
}
