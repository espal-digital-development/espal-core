// Code generated by espal-store-synthesizer. DO NOT EDIT.
package shipment

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ LineEntity = &Line{}

type LineEntity interface {
	database.Model
	Sorting() uint
	SetSorting(sorting uint)
	ShipmentID() string
	SetShipmentID(shipmentID string)
	SaleOrderLineID() *string
	SetSaleOrderLineID(saleOrderLineID *string)
	ProductVariantID() *string
	SetProductVariantID(productVariantID *string)
	BundledProductID() *string
	SetBundledProductID(bundledProductID *string)
	Quantity() int
	SetQuantity(quantity int)
	Comments() *string
	SetComments(comments *string)
}

// ID returns id.
func (line *Line) ID() string {
	return line.id
}

// CreatedByID returns createdByID.
func (line *Line) CreatedByID() string {
	return line.createdByID
}

// SetCreatedByID sets the createdByID.
func (line *Line) SetCreatedByID(createdByID string) {
	line.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (line *Line) UpdatedByID() *string {
	return line.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (line *Line) SetUpdatedByID(updatedByID *string) {
	line.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (line *Line) CreatedAt() time.Time {
	return line.createdAt
}

// SetCreatedAt sets the createdAt.
func (line *Line) SetCreatedAt(createdAt time.Time) {
	line.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (line *Line) UpdatedAt() *time.Time {
	return line.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (line *Line) SetUpdatedAt(updatedAt *time.Time) {
	line.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (line *Line) CreatedByFirstName() *string {
	return line.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (line *Line) SetCreatedByFirstName(createdByFirstName *string) {
	line.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (line *Line) CreatedBySurname() *string {
	return line.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (line *Line) SetCreatedBySurname(createdBySurname *string) {
	line.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (line *Line) UpdatedByFirstName() *string {
	return line.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (line *Line) SetUpdatedByFirstName(updatedByFirstName *string) {
	line.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (line *Line) UpdatedBySurname() *string {
	return line.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (line *Line) SetUpdatedBySurname(updatedBySurname *string) {
	line.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (line *Line) IsUpdated() bool {
	return line.updatedByID != nil
}

// Sorting returns sorting.
func (line *Line) Sorting() uint {
	return line.sorting
}

// SetSorting sets the sorting.
func (line *Line) SetSorting(sorting uint) {
	line.sorting = sorting
}

// ShipmentID returns shipmentID.
func (line *Line) ShipmentID() string {
	return line.shipmentID
}

// SetShipmentID sets the shipmentID.
func (line *Line) SetShipmentID(shipmentID string) {
	line.shipmentID = shipmentID
}

// SaleOrderLineID returns saleOrderLineID.
func (line *Line) SaleOrderLineID() *string {
	return line.saleOrderLineID
}

// SetSaleOrderLineID sets the saleOrderLineID.
func (line *Line) SetSaleOrderLineID(saleOrderLineID *string) {
	line.saleOrderLineID = saleOrderLineID
}

// ProductVariantID returns productVariantID.
func (line *Line) ProductVariantID() *string {
	return line.productVariantID
}

// SetProductVariantID sets the productVariantID.
func (line *Line) SetProductVariantID(productVariantID *string) {
	line.productVariantID = productVariantID
}

// BundledProductID returns bundledProductID.
func (line *Line) BundledProductID() *string {
	return line.bundledProductID
}

// SetBundledProductID sets the bundledProductID.
func (line *Line) SetBundledProductID(bundledProductID *string) {
	line.bundledProductID = bundledProductID
}

// Quantity returns quantity.
func (line *Line) Quantity() int {
	return line.quantity
}

// SetQuantity sets the quantity.
func (line *Line) SetQuantity(quantity int) {
	line.quantity = quantity
}

// Comments returns comments.
func (line *Line) Comments() *string {
	return line.comments
}

// SetComments sets the comments.
func (line *Line) SetComments(comments *string) {
	line.comments = comments
}

func newLine() *Line {
	return &Line{}
}

// New returns a new instance of LineEntity.
func NewLineEntity() LineEntity {
	return newLine()
}
