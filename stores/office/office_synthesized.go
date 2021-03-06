// Code generated by espal-store-synthesizer. DO NOT EDIT.
package office

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ OfficeEntity = &Office{}

type OfficeEntity interface {
	database.Model
	Active() bool
	SetActive(active bool)
	Sorting() uint
	SetSorting(sorting uint)
	PrimaryContactPerson() *string
	SetPrimaryContactPerson(primaryContactPerson *string)
	Street() string
	SetStreet(street string)
	StreetLine2() *string
	SetStreetLine2(streetLine2 *string)
	Number() string
	SetNumber(number string)
	NumberAddition() *string
	SetNumberAddition(numberAddition *string)
	ZipCode() string
	SetZipCode(zipCode string)
	City() string
	SetCity(city string)
	State() *string
	SetState(state *string)
	Country() *uint16
	SetCountry(country *uint16)
	PhoneNumber() *string
	SetPhoneNumber(phoneNumber *string)
	Email() *string
	SetEmail(email *string)
}

// ID returns id.
func (o *Office) ID() string {
	return o.id
}

// CreatedByID returns createdByID.
func (o *Office) CreatedByID() string {
	return o.createdByID
}

// SetCreatedByID sets the createdByID.
func (o *Office) SetCreatedByID(createdByID string) {
	o.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (o *Office) UpdatedByID() *string {
	return o.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (o *Office) SetUpdatedByID(updatedByID *string) {
	o.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (o *Office) CreatedAt() time.Time {
	return o.createdAt
}

// SetCreatedAt sets the createdAt.
func (o *Office) SetCreatedAt(createdAt time.Time) {
	o.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (o *Office) UpdatedAt() *time.Time {
	return o.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (o *Office) SetUpdatedAt(updatedAt *time.Time) {
	o.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (o *Office) CreatedByFirstName() *string {
	return o.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (o *Office) SetCreatedByFirstName(createdByFirstName *string) {
	o.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (o *Office) CreatedBySurname() *string {
	return o.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (o *Office) SetCreatedBySurname(createdBySurname *string) {
	o.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (o *Office) UpdatedByFirstName() *string {
	return o.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (o *Office) SetUpdatedByFirstName(updatedByFirstName *string) {
	o.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (o *Office) UpdatedBySurname() *string {
	return o.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (o *Office) SetUpdatedBySurname(updatedBySurname *string) {
	o.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (o *Office) IsUpdated() bool {
	return o.updatedByID != nil
}

// Active returns active.
func (o *Office) Active() bool {
	return o.active
}

// SetActive sets the active.
func (o *Office) SetActive(active bool) {
	o.active = active
}

// Sorting returns sorting.
func (o *Office) Sorting() uint {
	return o.sorting
}

// SetSorting sets the sorting.
func (o *Office) SetSorting(sorting uint) {
	o.sorting = sorting
}

// PrimaryContactPerson returns primaryContactPerson.
func (o *Office) PrimaryContactPerson() *string {
	return o.primaryContactPerson
}

// SetPrimaryContactPerson sets the primaryContactPerson.
func (o *Office) SetPrimaryContactPerson(primaryContactPerson *string) {
	o.primaryContactPerson = primaryContactPerson
}

// Street returns street.
func (o *Office) Street() string {
	return o.street
}

// SetStreet sets the street.
func (o *Office) SetStreet(street string) {
	o.street = street
}

// StreetLine2 returns streetLine2.
func (o *Office) StreetLine2() *string {
	return o.streetLine2
}

// SetStreetLine2 sets the streetLine2.
func (o *Office) SetStreetLine2(streetLine2 *string) {
	o.streetLine2 = streetLine2
}

// Number returns number.
func (o *Office) Number() string {
	return o.number
}

// SetNumber sets the number.
func (o *Office) SetNumber(number string) {
	o.number = number
}

// NumberAddition returns numberAddition.
func (o *Office) NumberAddition() *string {
	return o.numberAddition
}

// SetNumberAddition sets the numberAddition.
func (o *Office) SetNumberAddition(numberAddition *string) {
	o.numberAddition = numberAddition
}

// ZipCode returns zipCode.
func (o *Office) ZipCode() string {
	return o.zipCode
}

// SetZipCode sets the zipCode.
func (o *Office) SetZipCode(zipCode string) {
	o.zipCode = zipCode
}

// City returns city.
func (o *Office) City() string {
	return o.city
}

// SetCity sets the city.
func (o *Office) SetCity(city string) {
	o.city = city
}

// State returns state.
func (o *Office) State() *string {
	return o.state
}

// SetState sets the state.
func (o *Office) SetState(state *string) {
	o.state = state
}

// Country returns country.
func (o *Office) Country() *uint16 {
	return o.country
}

// SetCountry sets the country.
func (o *Office) SetCountry(country *uint16) {
	o.country = country
}

// PhoneNumber returns phoneNumber.
func (o *Office) PhoneNumber() *string {
	return o.phoneNumber
}

// SetPhoneNumber sets the phoneNumber.
func (o *Office) SetPhoneNumber(phoneNumber *string) {
	o.phoneNumber = phoneNumber
}

// Email returns email.
func (o *Office) Email() *string {
	return o.email
}

// SetEmail sets the email.
func (o *Office) SetEmail(email *string) {
	o.email = email
}

func newOffice() *Office {
	return &Office{}
}

// New returns a new instance of OfficeEntity.
func NewOfficeEntity() OfficeEntity {
	return newOffice()
}
