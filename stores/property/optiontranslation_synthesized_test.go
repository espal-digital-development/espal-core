// Code generated by espal-store-synthesizer. DO NOT EDIT.
package property_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/property"
)

func TestOptionTranslationTable(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	if o.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestOptionTranslationTableAlias(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	if o.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestOptionTranslationIsUpdated(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	o.IsUpdated()
}

func TestOptionTranslationID(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	o.ID()
}

func TestOptionTranslationCreatedByID(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetCreatedByID(testValue)
	if testValue != o.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationUpdatedByID(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetUpdatedByID(&testValue)
	if &testValue != o.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationCreatedAt(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := time.Now()
	o.SetCreatedAt(testValue)
	if testValue != o.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationUpdatedAt(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := time.Now()
	o.SetUpdatedAt(&testValue)
	if &testValue != o.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationCreatedByFirstName(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetCreatedByFirstName(&testValue)
	if &testValue != o.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationCreatedBySurname(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetCreatedBySurname(&testValue)
	if &testValue != o.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationUpdatedByFirstName(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetUpdatedByFirstName(&testValue)
	if &testValue != o.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationUpdatedBySurname(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetUpdatedBySurname(&testValue)
	if &testValue != o.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationLanguage(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := uint16(65000)
	o.SetLanguage(testValue)
	if testValue != o.Language() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationField(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := uint16(65000)
	o.SetField(testValue)
	if testValue != o.Field() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationValue(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetValue(testValue)
	if testValue != o.Value() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionTranslationOptionID(t *testing.T) {
	o := property.NewOptionTranslationEntity()
	testValue := "testValue"
	o.SetOptionID(testValue)
	if testValue != o.OptionID() {
		t.Fatal("Getter did not return the Set value")
	}
}
