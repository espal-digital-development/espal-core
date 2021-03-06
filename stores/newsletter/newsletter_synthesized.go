// Code generated by espal-store-synthesizer. DO NOT EDIT.
package newsletter

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ NewsletterEntity = &Newsletter{}

type NewsletterEntity interface {
	database.Model
	DomainID() string
	SetDomainID(domainID string)
	Active() bool
	SetActive(active bool)
	SendAt() *time.Time
	SetSendAt(sendAt *time.Time)
}

// TableName returns the table name that belongs to the current model.
func (n *Newsletter) TableName() string {
	return "Newsletter"
}

// TableAlias returns the unique resolved table alias for use in queries.
func (n *Newsletter) TableAlias() string {
	return "ne"
}

// ID returns id.
func (n *Newsletter) ID() string {
	return n.id
}

// CreatedByID returns createdByID.
func (n *Newsletter) CreatedByID() string {
	return n.createdByID
}

// SetCreatedByID sets the createdByID.
func (n *Newsletter) SetCreatedByID(createdByID string) {
	n.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (n *Newsletter) UpdatedByID() *string {
	return n.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (n *Newsletter) SetUpdatedByID(updatedByID *string) {
	n.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (n *Newsletter) CreatedAt() time.Time {
	return n.createdAt
}

// SetCreatedAt sets the createdAt.
func (n *Newsletter) SetCreatedAt(createdAt time.Time) {
	n.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (n *Newsletter) UpdatedAt() *time.Time {
	return n.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (n *Newsletter) SetUpdatedAt(updatedAt *time.Time) {
	n.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (n *Newsletter) CreatedByFirstName() *string {
	return n.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (n *Newsletter) SetCreatedByFirstName(createdByFirstName *string) {
	n.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (n *Newsletter) CreatedBySurname() *string {
	return n.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (n *Newsletter) SetCreatedBySurname(createdBySurname *string) {
	n.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (n *Newsletter) UpdatedByFirstName() *string {
	return n.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (n *Newsletter) SetUpdatedByFirstName(updatedByFirstName *string) {
	n.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (n *Newsletter) UpdatedBySurname() *string {
	return n.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (n *Newsletter) SetUpdatedBySurname(updatedBySurname *string) {
	n.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (n *Newsletter) IsUpdated() bool {
	return n.updatedByID != nil
}

// DomainID returns domainID.
func (n *Newsletter) DomainID() string {
	return n.domainID
}

// SetDomainID sets the domainID.
func (n *Newsletter) SetDomainID(domainID string) {
	n.domainID = domainID
}

// Active returns active.
func (n *Newsletter) Active() bool {
	return n.active
}

// SetActive sets the active.
func (n *Newsletter) SetActive(active bool) {
	n.active = active
}

// SendAt returns sendAt.
func (n *Newsletter) SendAt() *time.Time {
	return n.sendAt
}

// SetSendAt sets the sendAt.
func (n *Newsletter) SetSendAt(sendAt *time.Time) {
	n.sendAt = sendAt
}

func newNewsletter() *Newsletter {
	return &Newsletter{}
}

// New returns a new instance of NewsletterEntity.
func NewNewsletterEntity() NewsletterEntity {
	return newNewsletter()
}
