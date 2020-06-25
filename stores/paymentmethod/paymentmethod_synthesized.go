// Code generated by espal-store-synthesizer. DO NOT EDIT.
package paymentmethod

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ PaymentMethodEntity = &PaymentMethod{}

type PaymentMethodEntity interface {
	database.Model
	Name() string
	SetName(name string)
	Description() *string
	SetDescription(description *string)
}

// TableName returns the table name that belongs to the current model.
func (p *PaymentMethod) TableName() string {
	return "PaymentMethod"
}

// ID returns id.
func (p *PaymentMethod) ID() string {
	return p.id
}

// CreatedByID returns createdByID.
func (p *PaymentMethod) CreatedByID() string {
	return p.createdByID
}

// SetCreatedByID sets the createdByID.
func (p *PaymentMethod) SetCreatedByID(createdByID string) {
	p.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (p *PaymentMethod) UpdatedByID() *string {
	return p.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (p *PaymentMethod) SetUpdatedByID(updatedByID *string) {
	p.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (p *PaymentMethod) CreatedAt() time.Time {
	return p.createdAt
}

// SetCreatedAt sets the createdAt.
func (p *PaymentMethod) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (p *PaymentMethod) UpdatedAt() *time.Time {
	return p.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (p *PaymentMethod) SetUpdatedAt(updatedAt *time.Time) {
	p.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (p *PaymentMethod) CreatedByFirstName() *string {
	return p.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (p *PaymentMethod) SetCreatedByFirstName(createdByFirstName *string) {
	p.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (p *PaymentMethod) CreatedBySurname() *string {
	return p.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (p *PaymentMethod) SetCreatedBySurname(createdBySurname *string) {
	p.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (p *PaymentMethod) UpdatedByFirstName() *string {
	return p.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (p *PaymentMethod) SetUpdatedByFirstName(updatedByFirstName *string) {
	p.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (p *PaymentMethod) UpdatedBySurname() *string {
	return p.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (p *PaymentMethod) SetUpdatedBySurname(updatedBySurname *string) {
	p.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (p *PaymentMethod) IsUpdated() bool {
	return p.updatedByID != nil
}

// Name returns name.
func (p *PaymentMethod) Name() string {
	return p.name
}

// SetName sets the name.
func (p *PaymentMethod) SetName(name string) {
	p.name = name
}

// Description returns description.
func (p *PaymentMethod) Description() *string {
	return p.description
}

// SetDescription sets the description.
func (p *PaymentMethod) SetDescription(description *string) {
	p.description = description
}

func newPaymentMethod() *PaymentMethod {
	return &PaymentMethod{}
}

// New returns a new instance of PaymentMethodEntity.
func NewPaymentMethodEntity() PaymentMethodEntity {
	return newPaymentMethod()
}
