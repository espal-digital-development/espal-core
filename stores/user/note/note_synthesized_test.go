// Code generated by espal-store-synthesizer. DO NOT EDIT.
package note_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/user/note"
)

func TestNoteTable(t *testing.T) {
	n := note.NewNoteEntity()
	if n.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestNoteTableAlias(t *testing.T) {
	n := note.NewNoteEntity()
	if n.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestNoteIsUpdated(t *testing.T) {
	n := note.NewNoteEntity()
	n.IsUpdated()
}

func TestNoteID(t *testing.T) {
	n := note.NewNoteEntity()
	n.ID()
}

func TestNoteCreatedByID(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetCreatedByID(testValue)
	if testValue != n.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteUpdatedByID(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetUpdatedByID(&testValue)
	if &testValue != n.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteCreatedAt(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := time.Now()
	n.SetCreatedAt(testValue)
	if testValue != n.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteUpdatedAt(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := time.Now()
	n.SetUpdatedAt(&testValue)
	if &testValue != n.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteCreatedByFirstName(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetCreatedByFirstName(&testValue)
	if &testValue != n.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteCreatedBySurname(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetCreatedBySurname(&testValue)
	if &testValue != n.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteUpdatedByFirstName(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetUpdatedByFirstName(&testValue)
	if &testValue != n.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteUpdatedBySurname(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetUpdatedBySurname(&testValue)
	if &testValue != n.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteUserID(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetUserID(testValue)
	if testValue != n.UserID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteTitle(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetTitle(&testValue)
	if &testValue != n.Title() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestNoteContents(t *testing.T) {
	n := note.NewNoteEntity()
	testValue := "testValue"
	n.SetContents(testValue)
	if testValue != n.Contents() {
		t.Fatal("Getter did not return the Set value")
	}
}
