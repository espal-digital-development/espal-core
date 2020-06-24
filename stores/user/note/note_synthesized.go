// Code generated by espal-store-synthesizer. DO NOT EDIT.
package note

import (
	"time"

	"github.com/espal-digital-development/espal-core/database"
)

var _ NoteEntity = &Note{}

type NoteEntity interface {
	database.Model
	UserID() string
	SetUserID(userID string)
	Title() *string
	SetTitle(title *string)
	Contents() string
	SetContents(contents string)
}

// ID returns id.
func (note *Note) ID() string {
	return note.id
}

// CreatedByID returns createdByID.
func (note *Note) CreatedByID() string {
	return note.createdByID
}

// SetCreatedByID sets the createdByID.
func (note *Note) SetCreatedByID(createdByID string) {
	note.createdByID = createdByID
}

// UpdatedByID returns updatedByID.
func (note *Note) UpdatedByID() *string {
	return note.updatedByID
}

// SetUpdatedByID sets the updatedByID.
func (note *Note) SetUpdatedByID(updatedByID *string) {
	note.updatedByID = updatedByID
}

// CreatedAt returns createdAt.
func (note *Note) CreatedAt() time.Time {
	return note.createdAt
}

// SetCreatedAt sets the createdAt.
func (note *Note) SetCreatedAt(createdAt time.Time) {
	note.createdAt = createdAt
}

// UpdatedAt returns updatedAt.
func (note *Note) UpdatedAt() *time.Time {
	return note.updatedAt
}

// SetUpdatedAt sets the updatedAt.
func (note *Note) SetUpdatedAt(updatedAt *time.Time) {
	note.updatedAt = updatedAt
}

// CreatedByFirstName returns createdByFirstName.
func (note *Note) CreatedByFirstName() *string {
	return note.createdByFirstName
}

// SetCreatedByFirstName sets the createdByFirstName.
func (note *Note) SetCreatedByFirstName(createdByFirstName *string) {
	note.createdByFirstName = createdByFirstName
}

// CreatedBySurname returns createdBySurname.
func (note *Note) CreatedBySurname() *string {
	return note.createdBySurname
}

// SetCreatedBySurname sets the createdBySurname.
func (note *Note) SetCreatedBySurname(createdBySurname *string) {
	note.createdBySurname = createdBySurname
}

// UpdatedByFirstName returns updatedByFirstName.
func (note *Note) UpdatedByFirstName() *string {
	return note.updatedByFirstName
}

// SetUpdatedByFirstName sets the updatedByFirstName.
func (note *Note) SetUpdatedByFirstName(updatedByFirstName *string) {
	note.updatedByFirstName = updatedByFirstName
}

// UpdatedBySurname returns updatedBySurname.
func (note *Note) UpdatedBySurname() *string {
	return note.updatedBySurname
}

// SetUpdatedBySurname sets the updatedBySurname.
func (note *Note) SetUpdatedBySurname(updatedBySurname *string) {
	note.updatedBySurname = updatedBySurname
}

// IsUpdated returns true if UpdatedByID is set.
func (note *Note) IsUpdated() bool {
	return note.updatedByID != nil
}

// UserID returns userID.
func (note *Note) UserID() string {
	return note.userID
}

// SetUserID sets the userID.
func (note *Note) SetUserID(userID string) {
	note.userID = userID
}

// Title returns title.
func (note *Note) Title() *string {
	return note.title
}

// SetTitle sets the title.
func (note *Note) SetTitle(title *string) {
	note.title = title
}

// Contents returns contents.
func (note *Note) Contents() string {
	return note.contents
}

// SetContents sets the contents.
func (note *Note) SetContents(contents string) {
	note.contents = contents
}

func newNote() *Note {
	return &Note{}
}

// New returns a new instance of NoteEntity.
func NewNoteEntity() NoteEntity {
	return newNote()
}
