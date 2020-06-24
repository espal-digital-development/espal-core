// Code generated by espal-store-synthesizer. DO NOT EDIT.
package blogpost_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/blogpost"
)

func TestTranslationTable(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	if translation.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestTranslationTableAlias(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	if translation.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestTranslationIsUpdated(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	translation.IsUpdated()
}

func TestTranslationID(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	translation.ID()
}

func TestTranslationCreatedByID(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetCreatedByID(testValue)
	if testValue != translation.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedByID(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetUpdatedByID(&testValue)
	if &testValue != translation.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationCreatedAt(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := time.Now()
	translation.SetCreatedAt(testValue)
	if testValue != translation.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedAt(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := time.Now()
	translation.SetUpdatedAt(&testValue)
	if &testValue != translation.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationCreatedByFirstName(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetCreatedByFirstName(&testValue)
	if &testValue != translation.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationCreatedBySurname(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetCreatedBySurname(&testValue)
	if &testValue != translation.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedByFirstName(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetUpdatedByFirstName(&testValue)
	if &testValue != translation.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationUpdatedBySurname(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetUpdatedBySurname(&testValue)
	if &testValue != translation.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationLanguage(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := uint16(65000)
	translation.SetLanguage(testValue)
	if testValue != translation.Language() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationField(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := uint16(65000)
	translation.SetField(testValue)
	if testValue != translation.Field() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationValue(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetValue(testValue)
	if testValue != translation.Value() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestTranslationBlogPostID(t *testing.T) {
	translation := blogpost.NewTranslationEntity()
	testValue := "testValue"
	translation.SetBlogPostID(testValue)
	if testValue != translation.BlogPostID() {
		t.Fatal("Getter did not return the Set value")
	}
}
