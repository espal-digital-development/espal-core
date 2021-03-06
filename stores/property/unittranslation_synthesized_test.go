// Code generated by espal-store-synthesizer. DO NOT EDIT.
package property_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/property"
)

func TestUnitTranslationTable(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	if u.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestUnitTranslationTableAlias(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	if u.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestUnitTranslationIsUpdated(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	u.IsUpdated()
}

func TestUnitTranslationID(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	u.ID()
}

func TestUnitTranslationCreatedByID(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetCreatedByID(testValue)
	if testValue != u.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationUpdatedByID(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetUpdatedByID(&testValue)
	if &testValue != u.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationCreatedAt(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := time.Now()
	u.SetCreatedAt(testValue)
	if testValue != u.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationUpdatedAt(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := time.Now()
	u.SetUpdatedAt(&testValue)
	if &testValue != u.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationCreatedByFirstName(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetCreatedByFirstName(&testValue)
	if &testValue != u.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationCreatedBySurname(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetCreatedBySurname(&testValue)
	if &testValue != u.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationUpdatedByFirstName(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetUpdatedByFirstName(&testValue)
	if &testValue != u.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationUpdatedBySurname(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetUpdatedBySurname(&testValue)
	if &testValue != u.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationLanguage(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := uint16(65000)
	u.SetLanguage(testValue)
	if testValue != u.Language() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationField(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := uint16(65000)
	u.SetField(testValue)
	if testValue != u.Field() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationUnitID(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetUnitID(testValue)
	if testValue != u.UnitID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationValue(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetValue(testValue)
	if testValue != u.Value() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestUnitTranslationDisplay(t *testing.T) {
	u := property.NewUnitTranslationEntity()
	testValue := "testValue"
	u.SetDisplay(&testValue)
	if &testValue != u.Display() {
		t.Fatal("Getter did not return the Set value")
	}
}
