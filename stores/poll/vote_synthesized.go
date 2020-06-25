// Code generated by espal-store-synthesizer. DO NOT EDIT.
package poll

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ VoteEntity = &Vote{}

type VoteEntity interface {
	database.Model
	PollOptionID() string
	SetPollOptionID(pollOptionID string)
}

// ID returns id.
func (v *Vote) ID() string {
	return v.id
}

// CreatedByID returns createdByID.
func (v *Vote) CreatedByID() string {
	return v.createdByID
}

// SetCreatedByID sets the createdByID.
func (v *Vote) SetCreatedByID(createdByID string) {
	v.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (v *Vote) UpdatedByID() *string {
	return v.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (v *Vote) SetUpdatedByID(updatedByID *string) {
	v.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (v *Vote) CreatedAt() time.Time {
	return v.createdAt
}

// SetCreatedAt sets the createdAt.
func (v *Vote) SetCreatedAt(createdAt time.Time) {
	v.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (v *Vote) UpdatedAt() *time.Time {
	return v.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (v *Vote) SetUpdatedAt(updatedAt *time.Time) {
	v.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (v *Vote) CreatedByFirstName() *string {
	return v.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (v *Vote) SetCreatedByFirstName(createdByFirstName *string) {
	v.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (v *Vote) CreatedBySurname() *string {
	return v.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (v *Vote) SetCreatedBySurname(createdBySurname *string) {
	v.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (v *Vote) UpdatedByFirstName() *string {
	return v.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (v *Vote) SetUpdatedByFirstName(updatedByFirstName *string) {
	v.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (v *Vote) UpdatedBySurname() *string {
	return v.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (v *Vote) SetUpdatedBySurname(updatedBySurname *string) {
	v.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (v *Vote) IsUpdated() bool {
	return v.updatedByID != nil
}

// PollOptionID returns pollOptionID.
func (v *Vote) PollOptionID() string {
	return v.pollOptionID
}

// SetPollOptionID sets the pollOptionID.
func (v *Vote) SetPollOptionID(pollOptionID string) {
	v.pollOptionID = pollOptionID
}

func newVote() *Vote {
	return &Vote{}
}

// New returns a new instance of VoteEntity.
func NewVoteEntity() VoteEntity {
	return newVote()
}
