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
func (l *Line) ID() string {
	return l.id
}

// CreatedByID returns createdByID.
func (l *Line) CreatedByID() string {
	return l.createdByID
}

// SetCreatedByID sets the createdByID.
func (l *Line) SetCreatedByID(createdByID string) {
	l.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (l *Line) UpdatedByID() *string {
	return l.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (l *Line) SetUpdatedByID(updatedByID *string) {
	l.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (l *Line) CreatedAt() time.Time {
	return l.createdAt
}

// SetCreatedAt sets the createdAt.
func (l *Line) SetCreatedAt(createdAt time.Time) {
	l.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (l *Line) UpdatedAt() *time.Time {
	return l.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (l *Line) SetUpdatedAt(updatedAt *time.Time) {
	l.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (l *Line) CreatedByFirstName() *string {
	return l.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (l *Line) SetCreatedByFirstName(createdByFirstName *string) {
	l.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (l *Line) CreatedBySurname() *string {
	return l.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (l *Line) SetCreatedBySurname(createdBySurname *string) {
	l.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (l *Line) UpdatedByFirstName() *string {
	return l.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (l *Line) SetUpdatedByFirstName(updatedByFirstName *string) {
	l.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (l *Line) UpdatedBySurname() *string {
	return l.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (l *Line) SetUpdatedBySurname(updatedBySurname *string) {
	l.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (l *Line) IsUpdated() bool {
	return l.updatedByID != nil
}

// Sorting returns sorting.
func (l *Line) Sorting() uint {
	return l.sorting
}

// SetSorting sets the sorting.
func (l *Line) SetSorting(sorting uint) {
	l.sorting = sorting
}

// ShipmentID returns shipmentID.
func (l *Line) ShipmentID() string {
	return l.shipmentID
}

// SetShipmentID sets the shipmentID.
func (l *Line) SetShipmentID(shipmentID string) {
	l.shipmentID = shipmentID
}

// SaleOrderLineID returns saleOrderLineID.
func (l *Line) SaleOrderLineID() *string {
	return l.saleOrderLineID
}

// SetSaleOrderLineID sets the saleOrderLineID.
func (l *Line) SetSaleOrderLineID(saleOrderLineID *string) {
	l.saleOrderLineID = saleOrderLineID
}

// ProductVariantID returns productVariantID.
func (l *Line) ProductVariantID() *string {
	return l.productVariantID
}

// SetProductVariantID sets the productVariantID.
func (l *Line) SetProductVariantID(productVariantID *string) {
	l.productVariantID = productVariantID
}

// BundledProductID returns bundledProductID.
func (l *Line) BundledProductID() *string {
	return l.bundledProductID
}

// SetBundledProductID sets the bundledProductID.
func (l *Line) SetBundledProductID(bundledProductID *string) {
	l.bundledProductID = bundledProductID
}

// Quantity returns quantity.
func (l *Line) Quantity() int {
	return l.quantity
}

// SetQuantity sets the quantity.
func (l *Line) SetQuantity(quantity int) {
	l.quantity = quantity
}

// Comments returns comments.
func (l *Line) Comments() *string {
	return l.comments
}

// SetComments sets the comments.
func (l *Line) SetComments(comments *string) {
	l.comments = comments
}

func newLine() *Line {
	return &Line{}
}

// New returns a new instance of LineEntity.
func NewLineEntity() LineEntity {
	return newLine()
}
