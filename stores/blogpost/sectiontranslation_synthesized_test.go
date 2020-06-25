// Code generated by espal-store-synthesizer. DO NOT EDIT.
package blogpost_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/blogpost"
)

func TestSectionTranslationTable(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	if s.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestSectionTranslationTableAlias(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	if s.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestSectionTranslationIsUpdated(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	s.IsUpdated()
}

func TestSectionTranslationID(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	s.ID()
}

func TestSectionTranslationCreatedByID(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetCreatedByID(testValue)
	if testValue != s.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationUpdatedByID(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetUpdatedByID(&testValue)
	if &testValue != s.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationCreatedAt(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := time.Now()
	s.SetCreatedAt(testValue)
	if testValue != s.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationUpdatedAt(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := time.Now()
	s.SetUpdatedAt(&testValue)
	if &testValue != s.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationCreatedByFirstName(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetCreatedByFirstName(&testValue)
	if &testValue != s.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationCreatedBySurname(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetCreatedBySurname(&testValue)
	if &testValue != s.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationUpdatedByFirstName(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetUpdatedByFirstName(&testValue)
	if &testValue != s.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationUpdatedBySurname(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetUpdatedBySurname(&testValue)
	if &testValue != s.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationLanguage(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := uint16(65000)
	s.SetLanguage(testValue)
	if testValue != s.Language() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationField(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := uint16(65000)
	s.SetField(testValue)
	if testValue != s.Field() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationValue(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetValue(testValue)
	if testValue != s.Value() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestSectionTranslationSectionID(t *testing.T) {
	s := blogpost.NewSectionTranslationEntity()
	testValue := "testValue"
	s.SetSectionID(testValue)
	if testValue != s.SectionID() {
		t.Fatal("Getter did not return the Set value")
	}
}
