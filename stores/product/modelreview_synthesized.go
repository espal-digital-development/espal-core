// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ ModelReviewEntity = &ModelReview{}

type ModelReviewEntity interface {
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
	ModelID() string
	SetModelID(modelID string)
}

// ID returns id.
func (modelReview *ModelReview) ID() string {
	return modelReview.id
}

// CreatedByID returns createdByID.
func (modelReview *ModelReview) CreatedByID() string {
	return modelReview.createdByID
}

// SetCreatedByID sets the createdByID.
func (modelReview *ModelReview) SetCreatedByID(createdByID string) {
	modelReview.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (modelReview *ModelReview) UpdatedByID() *string {
	return modelReview.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (modelReview *ModelReview) SetUpdatedByID(updatedByID *string) {
	modelReview.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (modelReview *ModelReview) CreatedAt() time.Time {
	return modelReview.createdAt
}

// SetCreatedAt sets the createdAt.
func (modelReview *ModelReview) SetCreatedAt(createdAt time.Time) {
	modelReview.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (modelReview *ModelReview) UpdatedAt() *time.Time {
	return modelReview.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (modelReview *ModelReview) SetUpdatedAt(updatedAt *time.Time) {
	modelReview.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (modelReview *ModelReview) CreatedByFirstName() *string {
	return modelReview.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (modelReview *ModelReview) SetCreatedByFirstName(createdByFirstName *string) {
	modelReview.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (modelReview *ModelReview) CreatedBySurname() *string {
	return modelReview.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (modelReview *ModelReview) SetCreatedBySurname(createdBySurname *string) {
	modelReview.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (modelReview *ModelReview) UpdatedByFirstName() *string {
	return modelReview.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (modelReview *ModelReview) SetUpdatedByFirstName(updatedByFirstName *string) {
	modelReview.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (modelReview *ModelReview) UpdatedBySurname() *string {
	return modelReview.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (modelReview *ModelReview) SetUpdatedBySurname(updatedBySurname *string) {
	modelReview.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (modelReview *ModelReview) IsUpdated() bool {
	return modelReview.updatedByID != nil
}

// ReviewedByID returns reviewedByID.
func (modelReview *ModelReview) ReviewedByID() *string {
	return modelReview.reviewedByID
}

// SetReviewedByID sets the reviewedByID.
func (modelReview *ModelReview) SetReviewedByID(reviewedByID *string) {
	modelReview.reviewedByID = reviewedByID
}

// ReviewedOnDate returns reviewedOnDate.
func (modelReview *ModelReview) ReviewedOnDate() *time.Time {
	return modelReview.reviewedOnDate
}

// SetReviewedOnDate sets the reviewedOnDate.
func (modelReview *ModelReview) SetReviewedOnDate(reviewedOnDate *time.Time) {
	modelReview.reviewedOnDate = reviewedOnDate
}

// ReviewNotes returns reviewNotes.
func (modelReview *ModelReview) ReviewNotes() *string {
	return modelReview.reviewNotes
}

// SetReviewNotes sets the reviewNotes.
func (modelReview *ModelReview) SetReviewNotes(reviewNotes *string) {
	modelReview.reviewNotes = reviewNotes
}

// Approved returns approved.
func (modelReview *ModelReview) Approved() *bool {
	return modelReview.approved
}

// SetApproved sets the approved.
func (modelReview *ModelReview) SetApproved(approved *bool) {
	modelReview.approved = approved
}

// Rating returns rating.
func (modelReview *ModelReview) Rating() float32 {
	return modelReview.rating
}

// SetRating sets the rating.
func (modelReview *ModelReview) SetRating(rating float32) {
	modelReview.rating = rating
}

// Title returns title.
func (modelReview *ModelReview) Title() string {
	return modelReview.title
}

// SetTitle sets the title.
func (modelReview *ModelReview) SetTitle(title string) {
	modelReview.title = title
}

// Description returns description.
func (modelReview *ModelReview) Description() string {
	return modelReview.description
}

// SetDescription sets the description.
func (modelReview *ModelReview) SetDescription(description string) {
	modelReview.description = description
}

// ModelID returns modelID.
func (modelReview *ModelReview) ModelID() string {
	return modelReview.modelID
}

// SetModelID sets the modelID.
func (modelReview *ModelReview) SetModelID(modelID string) {
	modelReview.modelID = modelID
}

func newModelReview() *ModelReview {
	return &ModelReview{}
}

// New returns a new instance of ModelReviewEntity.
func NewModelReviewEntity() ModelReviewEntity {
	return newModelReview()
}
