// Code generated by espal-store-synthesizer. DO NOT EDIT.
package invoice

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ InvoiceEntity = &Invoice{}

type InvoiceEntity interface {
	database.Model
	DomainID() string
	SetDomainID(domainID string)
	UserID() string
	SetUserID(userID string)
	SaleOrderID() string
	SetSaleOrderID(saleOrderID string)
	Currency() uint
	SetCurrency(currency uint)
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
	Comments() *string
	SetComments(comments *string)
	SellingPartyAutograph() *string
	SetSellingPartyAutograph(sellingPartyAutograph *string)
	BuyingPartyAutograph() *string
	SetBuyingPartyAutograph(buyingPartyAutograph *string)
}

// TableName returns the table name that belongs to the current model.
func (i *Invoice) TableName() string {
	return "Invoice"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (i *Invoice) TableAlias() string {
	return "ie"
}

// ID returns id.
func (i *Invoice) ID() string {
	return i.id
}

// CreatedByID returns createdByID.
func (i *Invoice) CreatedByID() string {
	return i.createdByID
}

// SetCreatedByID sets the createdByID.
func (i *Invoice) SetCreatedByID(createdByID string) {
	i.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (i *Invoice) UpdatedByID() *string {
	return i.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (i *Invoice) SetUpdatedByID(updatedByID *string) {
	i.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (i *Invoice) CreatedAt() time.Time {
	return i.createdAt
}

// SetCreatedAt sets the createdAt.
func (i *Invoice) SetCreatedAt(createdAt time.Time) {
	i.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (i *Invoice) UpdatedAt() *time.Time {
	return i.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (i *Invoice) SetUpdatedAt(updatedAt *time.Time) {
	i.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (i *Invoice) CreatedByFirstName() *string {
	return i.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (i *Invoice) SetCreatedByFirstName(createdByFirstName *string) {
	i.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (i *Invoice) CreatedBySurname() *string {
	return i.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (i *Invoice) SetCreatedBySurname(createdBySurname *string) {
	i.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (i *Invoice) UpdatedByFirstName() *string {
	return i.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (i *Invoice) SetUpdatedByFirstName(updatedByFirstName *string) {
	i.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (i *Invoice) UpdatedBySurname() *string {
	return i.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (i *Invoice) SetUpdatedBySurname(updatedBySurname *string) {
	i.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (i *Invoice) IsUpdated() bool {
	return i.updatedByID != nil
}

// DomainID returns domainID.
func (i *Invoice) DomainID() string {
	return i.domainID
}

// SetDomainID sets the domainID.
func (i *Invoice) SetDomainID(domainID string) {
	i.domainID = domainID
}

// UserID returns userID.
func (i *Invoice) UserID() string {
	return i.userID
}

// SetUserID sets the userID.
func (i *Invoice) SetUserID(userID string) {
	i.userID = userID
}

// SaleOrderID returns saleOrderID.
func (i *Invoice) SaleOrderID() string {
	return i.saleOrderID
}

// SetSaleOrderID sets the saleOrderID.
func (i *Invoice) SetSaleOrderID(saleOrderID string) {
	i.saleOrderID = saleOrderID
}

// Currency returns currency.
func (i *Invoice) Currency() uint {
	return i.currency
}

// SetCurrency sets the currency.
func (i *Invoice) SetCurrency(currency uint) {
	i.currency = currency
}

// Code returns code.
func (i *Invoice) Code() *string {
	return i.code
}

// SetCode sets the code.
func (i *Invoice) SetCode(code *string) {
	i.code = code
}

// UserInfoBusiness returns userInfoBusiness.
func (i *Invoice) UserInfoBusiness() bool {
	return i.userInfoBusiness
}

// SetUserInfoBusiness sets the userInfoBusiness.
func (i *Invoice) SetUserInfoBusiness(userInfoBusiness bool) {
	i.userInfoBusiness = userInfoBusiness
}

// UserInfoBusinessCocNumber returns userInfoBusinessCocNumber.
func (i *Invoice) UserInfoBusinessCocNumber() *string {
	return i.userInfoBusinessCocNumber
}

// SetUserInfoBusinessCocNumber sets the userInfoBusinessCocNumber.
func (i *Invoice) SetUserInfoBusinessCocNumber(userInfoBusinessCocNumber *string) {
	i.userInfoBusinessCocNumber = userInfoBusinessCocNumber
}

// UserInfoFirstName returns userInfoFirstName.
func (i *Invoice) UserInfoFirstName() string {
	return i.userInfoFirstName
}

// SetUserInfoFirstName sets the userInfoFirstName.
func (i *Invoice) SetUserInfoFirstName(userInfoFirstName string) {
	i.userInfoFirstName = userInfoFirstName
}

// UserInfoSurname returns userInfoSurname.
func (i *Invoice) UserInfoSurname() string {
	return i.userInfoSurname
}

// SetUserInfoSurname sets the userInfoSurname.
func (i *Invoice) SetUserInfoSurname(userInfoSurname string) {
	i.userInfoSurname = userInfoSurname
}

// UserInfoStreet returns userInfoStreet.
func (i *Invoice) UserInfoStreet() string {
	return i.userInfoStreet
}

// SetUserInfoStreet sets the userInfoStreet.
func (i *Invoice) SetUserInfoStreet(userInfoStreet string) {
	i.userInfoStreet = userInfoStreet
}

// UserInfoStreetLine2 returns userInfoStreetLine2.
func (i *Invoice) UserInfoStreetLine2() *string {
	return i.userInfoStreetLine2
}

// SetUserInfoStreetLine2 sets the userInfoStreetLine2.
func (i *Invoice) SetUserInfoStreetLine2(userInfoStreetLine2 *string) {
	i.userInfoStreetLine2 = userInfoStreetLine2
}

// UserInfoNumber returns userInfoNumber.
func (i *Invoice) UserInfoNumber() string {
	return i.userInfoNumber
}

// SetUserInfoNumber sets the userInfoNumber.
func (i *Invoice) SetUserInfoNumber(userInfoNumber string) {
	i.userInfoNumber = userInfoNumber
}

// UserInfoNumberAddition returns userInfoNumberAddition.
func (i *Invoice) UserInfoNumberAddition() *string {
	return i.userInfoNumberAddition
}

// SetUserInfoNumberAddition sets the userInfoNumberAddition.
func (i *Invoice) SetUserInfoNumberAddition(userInfoNumberAddition *string) {
	i.userInfoNumberAddition = userInfoNumberAddition
}

// UserInfoZipCode returns userInfoZipCode.
func (i *Invoice) UserInfoZipCode() string {
	return i.userInfoZipCode
}

// SetUserInfoZipCode sets the userInfoZipCode.
func (i *Invoice) SetUserInfoZipCode(userInfoZipCode string) {
	i.userInfoZipCode = userInfoZipCode
}

// UserInfoCity returns userInfoCity.
func (i *Invoice) UserInfoCity() string {
	return i.userInfoCity
}

// SetUserInfoCity sets the userInfoCity.
func (i *Invoice) SetUserInfoCity(userInfoCity string) {
	i.userInfoCity = userInfoCity
}

// UserInfoState returns userInfoState.
func (i *Invoice) UserInfoState() *uint {
	return i.userInfoState
}

// SetUserInfoState sets the userInfoState.
func (i *Invoice) SetUserInfoState(userInfoState *uint) {
	i.userInfoState = userInfoState
}

// UserInfoCountry returns userInfoCountry.
func (i *Invoice) UserInfoCountry() *uint16 {
	return i.userInfoCountry
}

// SetUserInfoCountry sets the userInfoCountry.
func (i *Invoice) SetUserInfoCountry(userInfoCountry *uint16) {
	i.userInfoCountry = userInfoCountry
}

// UserInfoPhoneNumber returns userInfoPhoneNumber.
func (i *Invoice) UserInfoPhoneNumber() *string {
	return i.userInfoPhoneNumber
}

// SetUserInfoPhoneNumber sets the userInfoPhoneNumber.
func (i *Invoice) SetUserInfoPhoneNumber(userInfoPhoneNumber *string) {
	i.userInfoPhoneNumber = userInfoPhoneNumber
}

// UserInfoEmail returns userInfoEmail.
func (i *Invoice) UserInfoEmail() *string {
	return i.userInfoEmail
}

// SetUserInfoEmail sets the userInfoEmail.
func (i *Invoice) SetUserInfoEmail(userInfoEmail *string) {
	i.userInfoEmail = userInfoEmail
}

// Comments returns comments.
func (i *Invoice) Comments() *string {
	return i.comments
}

// SetComments sets the comments.
func (i *Invoice) SetComments(comments *string) {
	i.comments = comments
}

// SellingPartyAutograph returns sellingPartyAutograph.
func (i *Invoice) SellingPartyAutograph() *string {
	return i.sellingPartyAutograph
}

// SetSellingPartyAutograph sets the sellingPartyAutograph.
func (i *Invoice) SetSellingPartyAutograph(sellingPartyAutograph *string) {
	i.sellingPartyAutograph = sellingPartyAutograph
}

// BuyingPartyAutograph returns buyingPartyAutograph.
func (i *Invoice) BuyingPartyAutograph() *string {
	return i.buyingPartyAutograph
}

// SetBuyingPartyAutograph sets the buyingPartyAutograph.
func (i *Invoice) SetBuyingPartyAutograph(buyingPartyAutograph *string) {
	i.buyingPartyAutograph = buyingPartyAutograph
}

func newInvoice() *Invoice {
	return &Invoice{}
}

// New returns a new instance of InvoiceEntity.
func NewInvoiceEntity() InvoiceEntity {
	return newInvoice()
}
