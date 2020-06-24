// Code generated by espal-store-synthesizer. DO NOT EDIT.
package task

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ TaskEntity = &Task{}

type TaskEntity interface {
	database.Model
	IssuedByID() string
	SetIssuedByID(issuedByID string)
	AssignedToID() *string
	SetAssignedToID(assignedToID *string)
	Description() string
	SetDescription(description string)
	CompletedNotes() *string
	SetCompletedNotes(completedNotes *string)
	CompletedAt() *time.Time
	SetCompletedAt(completedAt *time.Time)
}

// ID returns id.
func (task *Task) ID() string {
	return task.id
}

// CreatedByID returns createdByID.
func (task *Task) CreatedByID() string {
	return task.createdByID
}

// SetCreatedByID sets the createdByID.
func (task *Task) SetCreatedByID(createdByID string) {
	task.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (task *Task) UpdatedByID() *string {
	return task.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (task *Task) SetUpdatedByID(updatedByID *string) {
	task.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (task *Task) CreatedAt() time.Time {
	return task.createdAt
}

// SetCreatedAt sets the createdAt.
func (task *Task) SetCreatedAt(createdAt time.Time) {
	task.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (task *Task) UpdatedAt() *time.Time {
	return task.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (task *Task) SetUpdatedAt(updatedAt *time.Time) {
	task.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (task *Task) CreatedByFirstName() *string {
	return task.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (task *Task) SetCreatedByFirstName(createdByFirstName *string) {
	task.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (task *Task) CreatedBySurname() *string {
	return task.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (task *Task) SetCreatedBySurname(createdBySurname *string) {
	task.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (task *Task) UpdatedByFirstName() *string {
	return task.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (task *Task) SetUpdatedByFirstName(updatedByFirstName *string) {
	task.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (task *Task) UpdatedBySurname() *string {
	return task.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (task *Task) SetUpdatedBySurname(updatedBySurname *string) {
	task.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (task *Task) IsUpdated() bool {
	return task.updatedByID != nil
}

// IssuedByID returns issuedByID.
func (task *Task) IssuedByID() string {
	return task.issuedByID
}

// SetIssuedByID sets the issuedByID.
func (task *Task) SetIssuedByID(issuedByID string) {
	task.issuedByID = issuedByID
}

// AssignedToID returns assignedToID.
func (task *Task) AssignedToID() *string {
	return task.assignedToID
}

// SetAssignedToID sets the assignedToID.
func (task *Task) SetAssignedToID(assignedToID *string) {
	task.assignedToID = assignedToID
}

// Description returns description.
func (task *Task) Description() string {
	return task.description
}

// SetDescription sets the description.
func (task *Task) SetDescription(description string) {
	task.description = description
}

// CompletedNotes returns completedNotes.
func (task *Task) CompletedNotes() *string {
	return task.completedNotes
}

// SetCompletedNotes sets the completedNotes.
func (task *Task) SetCompletedNotes(completedNotes *string) {
	task.completedNotes = completedNotes
}

// CompletedAt returns completedAt.
func (task *Task) CompletedAt() *time.Time {
	return task.completedAt
}

// SetCompletedAt sets the completedAt.
func (task *Task) SetCompletedAt(completedAt *time.Time) {
	task.completedAt = completedAt
}

func newTask() *Task {
	return &Task{}
}

// New returns a new instance of TaskEntity.
func NewTaskEntity() TaskEntity {
	return newTask()
}
