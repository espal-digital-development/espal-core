// Code generated by espal-store-synthesizer. DO NOT EDIT.
package address

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ AddressEntity = &Address{}

type AddressEntity interface {
	database.Model
	UserID() string
	SetUserID(userID string)
	Active() bool
	SetActive(active bool)
	FirstName() *string
	SetFirstName(firstName *string)
	Surname() *string
	SetSurname(surname *string)
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
func (a *Address) ID() string {
	return a.id
}

// CreatedByID returns createdByID.
func (a *Address) CreatedByID() string {
	return a.createdByID
}

// SetCreatedByID sets the createdByID.
func (a *Address) SetCreatedByID(createdByID string) {
	a.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (a *Address) UpdatedByID() *string {
	return a.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (a *Address) SetUpdatedByID(updatedByID *string) {
	a.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (a *Address) CreatedAt() time.Time {
	return a.createdAt
}

// SetCreatedAt sets the createdAt.
func (a *Address) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (a *Address) UpdatedAt() *time.Time {
	return a.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (a *Address) SetUpdatedAt(updatedAt *time.Time) {
	a.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (a *Address) CreatedByFirstName() *string {
	return a.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (a *Address) SetCreatedByFirstName(createdByFirstName *string) {
	a.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (a *Address) CreatedBySurname() *string {
	return a.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (a *Address) SetCreatedBySurname(createdBySurname *string) {
	a.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (a *Address) UpdatedByFirstName() *string {
	return a.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (a *Address) SetUpdatedByFirstName(updatedByFirstName *string) {
	a.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (a *Address) UpdatedBySurname() *string {
	return a.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (a *Address) SetUpdatedBySurname(updatedBySurname *string) {
	a.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (a *Address) IsUpdated() bool {
	return a.updatedByID != nil
}

// UserID returns userID.
func (a *Address) UserID() string {
	return a.userID
}

// SetUserID sets the userID.
func (a *Address) SetUserID(userID string) {
	a.userID = userID
}

// Active returns active.
func (a *Address) Active() bool {
	return a.active
}

// SetActive sets the active.
func (a *Address) SetActive(active bool) {
	a.active = active
}

// FirstName returns firstName.
func (a *Address) FirstName() *string {
	return a.firstName
}

// SetFirstName sets the firstName.
func (a *Address) SetFirstName(firstName *string) {
	a.firstName = firstName
}

// Surname returns surname.
func (a *Address) Surname() *string {
	return a.surname
}

// SetSurname sets the surname.
func (a *Address) SetSurname(surname *string) {
	a.surname = surname
}

// Street returns street.
func (a *Address) Street() string {
	return a.street
}

// SetStreet sets the street.
func (a *Address) SetStreet(street string) {
	a.street = street
}

// StreetLine2 returns streetLine2.
func (a *Address) StreetLine2() *string {
	return a.streetLine2
}

// SetStreetLine2 sets the streetLine2.
func (a *Address) SetStreetLine2(streetLine2 *string) {
	a.streetLine2 = streetLine2
}

// Number returns number.
func (a *Address) Number() string {
	return a.number
}

// SetNumber sets the number.
func (a *Address) SetNumber(number string) {
	a.number = number
}

// NumberAddition returns numberAddition.
func (a *Address) NumberAddition() *string {
	return a.numberAddition
}

// SetNumberAddition sets the numberAddition.
func (a *Address) SetNumberAddition(numberAddition *string) {
	a.numberAddition = numberAddition
}

// ZipCode returns zipCode.
func (a *Address) ZipCode() string {
	return a.zipCode
}

// SetZipCode sets the zipCode.
func (a *Address) SetZipCode(zipCode string) {
	a.zipCode = zipCode
}

// City returns city.
func (a *Address) City() string {
	return a.city
}

// SetCity sets the city.
func (a *Address) SetCity(city string) {
	a.city = city
}

// State returns state.
func (a *Address) State() *string {
	return a.state
}

// SetState sets the state.
func (a *Address) SetState(state *string) {
	a.state = state
}

// Country returns country.
func (a *Address) Country() *uint16 {
	return a.country
}

// SetCountry sets the country.
func (a *Address) SetCountry(country *uint16) {
	a.country = country
}

// PhoneNumber returns phoneNumber.
func (a *Address) PhoneNumber() *string {
	return a.phoneNumber
}

// SetPhoneNumber sets the phoneNumber.
func (a *Address) SetPhoneNumber(phoneNumber *string) {
	a.phoneNumber = phoneNumber
}

// Email returns email.
func (a *Address) Email() *string {
	return a.email
}

// SetEmail sets the email.
func (a *Address) SetEmail(email *string) {
	a.email = email
}

func newAddress() *Address {
	return &Address{}
}

// New returns a new instance of AddressEntity.
func NewAddressEntity() AddressEntity {
	return newAddress()
}
