// Code generated by espal-store-synthesizer. DO NOT EDIT.
package price

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ PriceEntity = &Price{}

type PriceEntity interface {
	database.Model
	BundledProductID() *string
	SetBundledProductID(bundledProductID *string)
	ProductModelID() *string
	SetProductModelID(productModelID *string)
	ProductVariantID() *string
	SetProductVariantID(productVariantID *string)
	DomainID() string
	SetDomainID(domainID string)
	PriceGroupID() string
	SetPriceGroupID(priceGroupID string)
	TaxGroup() uint
	SetTaxGroup(taxGroup uint)
	Currency() uint16
	SetCurrency(currency uint16)
	Price() float32
	SetPrice(price float32)
}

// TableName returns the table name that belongs to the current model.
func (price *Price) TableName() string {
	return "Price"
}

// ID returns id.
func (price *Price) ID() string {
	return price.id
}

// CreatedByID returns createdByID.
func (price *Price) CreatedByID() string {
	return price.createdByID
}

// SetCreatedByID sets the createdByID.
func (price *Price) SetCreatedByID(createdByID string) {
	price.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (price *Price) UpdatedByID() *string {
	return price.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (price *Price) SetUpdatedByID(updatedByID *string) {
	price.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (price *Price) CreatedAt() time.Time {
	return price.createdAt
}

// SetCreatedAt sets the createdAt.
func (price *Price) SetCreatedAt(createdAt time.Time) {
	price.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (price *Price) UpdatedAt() *time.Time {
	return price.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (price *Price) SetUpdatedAt(updatedAt *time.Time) {
	price.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (price *Price) CreatedByFirstName() *string {
	return price.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (price *Price) SetCreatedByFirstName(createdByFirstName *string) {
	price.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (price *Price) CreatedBySurname() *string {
	return price.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (price *Price) SetCreatedBySurname(createdBySurname *string) {
	price.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (price *Price) UpdatedByFirstName() *string {
	return price.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (price *Price) SetUpdatedByFirstName(updatedByFirstName *string) {
	price.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (price *Price) UpdatedBySurname() *string {
	return price.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (price *Price) SetUpdatedBySurname(updatedBySurname *string) {
	price.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (price *Price) IsUpdated() bool {
	return price.updatedByID != nil
}

// BundledProductID returns bundledProductID.
func (price *Price) BundledProductID() *string {
	return price.bundledProductID
}

// SetBundledProductID sets the bundledProductID.
func (price *Price) SetBundledProductID(bundledProductID *string) {
	price.bundledProductID = bundledProductID
}

// ProductModelID returns productModelID.
func (price *Price) ProductModelID() *string {
	return price.productModelID
}

// SetProductModelID sets the productModelID.
func (price *Price) SetProductModelID(productModelID *string) {
	price.productModelID = productModelID
}

// ProductVariantID returns productVariantID.
func (price *Price) ProductVariantID() *string {
	return price.productVariantID
}

// SetProductVariantID sets the productVariantID.
func (price *Price) SetProductVariantID(productVariantID *string) {
	price.productVariantID = productVariantID
}

// DomainID returns domainID.
func (price *Price) DomainID() string {
	return price.domainID
}

// SetDomainID sets the domainID.
func (price *Price) SetDomainID(domainID string) {
	price.domainID = domainID
}

// PriceGroupID returns priceGroupID.
func (price *Price) PriceGroupID() string {
	return price.priceGroupID
}

// SetPriceGroupID sets the priceGroupID.
func (price *Price) SetPriceGroupID(priceGroupID string) {
	price.priceGroupID = priceGroupID
}

// TaxGroup returns taxGroup.
func (price *Price) TaxGroup() uint {
	return price.taxGroup
}

// SetTaxGroup sets the taxGroup.
func (price *Price) SetTaxGroup(taxGroup uint) {
	price.taxGroup = taxGroup
}

// Currency returns currency.
func (price *Price) Currency() uint16 {
	return price.currency
}

// SetCurrency sets the currency.
func (price *Price) SetCurrency(currency uint16) {
	price.currency = currency
}

// Price returns price.
func (price *Price) Price() float32 {
	return price.price
}

// SetPrice sets the price.
func (priceEntity *Price) SetPrice(price float32) {
	priceEntity.price = price
}

func newPrice() *Price {
	return &Price{}
}

// New returns a new instance of PriceEntity.
func NewPriceEntity() PriceEntity {
	return newPrice()
}
