// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ VariantReviewEntity = &VariantReview{}

type VariantReviewEntity interface {
	database.Model
	ReviewedByID() *string
	SetReviewedByID(reviewedByID *string)
	ReviewedOnDate() *time.Time
	SetReviewedOnDate(reviewedOnDate *time.Time)
	ReviewNotes() *string
	SetReviewNotes(reviewNotes *string)
	Approved() *bool
	SetApproved(approved *bool)
	Rating() float32
	SetRating(rating float32)
	Title() string
	SetTitle(title string)
	Description() string
	SetDescription(description string)
	VariantID() string
	SetVariantID(variantID string)
}

// ID returns id.
func (v *VariantReview) ID() string {
	return v.id
}

// CreatedByID returns createdByID.
func (v *VariantReview) CreatedByID() string {
	return v.createdByID
}

// SetCreatedByID sets the createdByID.
func (v *VariantReview) SetCreatedByID(createdByID string) {
	v.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (v *VariantReview) UpdatedByID() *string {
	return v.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (v *VariantReview) SetUpdatedByID(updatedByID *string) {
	v.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (v *VariantReview) CreatedAt() time.Time {
	return v.createdAt
}

// SetCreatedAt sets the createdAt.
func (v *VariantReview) SetCreatedAt(createdAt time.Time) {
	v.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (v *VariantReview) UpdatedAt() *time.Time {
	return v.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (v *VariantReview) SetUpdatedAt(updatedAt *time.Time) {
	v.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (v *VariantReview) CreatedByFirstName() *string {
	return v.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (v *VariantReview) SetCreatedByFirstName(createdByFirstName *string) {
	v.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (v *VariantReview) CreatedBySurname() *string {
	return v.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (v *VariantReview) SetCreatedBySurname(createdBySurname *string) {
	v.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (v *VariantReview) UpdatedByFirstName() *string {
	return v.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (v *VariantReview) SetUpdatedByFirstName(updatedByFirstName *string) {
	v.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (v *VariantReview) UpdatedBySurname() *string {
	return v.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (v *VariantReview) SetUpdatedBySurname(updatedBySurname *string) {
	v.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (v *VariantReview) IsUpdated() bool {
	return v.updatedByID != nil
}

// ReviewedByID returns reviewedByID.
func (v *VariantReview) ReviewedByID() *string {
	return v.reviewedByID
}

// SetReviewedByID sets the reviewedByID.
func (v *VariantReview) SetReviewedByID(reviewedByID *string) {
	v.reviewedByID = reviewedByID
}

// ReviewedOnDate returns reviewedOnDate.
func (v *VariantReview) ReviewedOnDate() *time.Time {
	return v.reviewedOnDate
}

// SetReviewedOnDate sets the reviewedOnDate.
func (v *VariantReview) SetReviewedOnDate(reviewedOnDate *time.Time) {
	v.reviewedOnDate = reviewedOnDate
}

// ReviewNotes returns reviewNotes.
func (v *VariantReview) ReviewNotes() *string {
	return v.reviewNotes
}

// SetReviewNotes sets the reviewNotes.
func (v *VariantReview) SetReviewNotes(reviewNotes *string) {
	v.reviewNotes = reviewNotes
}

// Approved returns approved.
func (v *VariantReview) Approved() *bool {
	return v.approved
}

// SetApproved sets the approved.
func (v *VariantReview) SetApproved(approved *bool) {
	v.approved = approved
}

// Rating returns rating.
func (v *VariantReview) Rating() float32 {
	return v.rating
}

// SetRating sets the rating.
func (v *VariantReview) SetRating(rating float32) {
	v.rating = rating
}

// Title returns title.
func (v *VariantReview) Title() string {
	return v.title
}

// SetTitle sets the title.
func (v *VariantReview) SetTitle(title string) {
	v.title = title
}

// Description returns description.
func (v *VariantReview) Description() string {
	return v.description
}

// SetDescription sets the description.
func (v *VariantReview) SetDescription(description string) {
	v.description = description
}

// VariantID returns variantID.
func (v *VariantReview) VariantID() string {
	return v.variantID
}

// SetVariantID sets the variantID.
func (v *VariantReview) SetVariantID(variantID string) {
	v.variantID = variantID
}

func newVariantReview() *VariantReview {
	return &VariantReview{}
}

// New returns a new instance of VariantReviewEntity.
func NewVariantReviewEntity() VariantReviewEntity {
	return newVariantReview()
}
