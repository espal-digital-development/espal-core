// Code generated by espal-store-synthesizer. DO NOT EDIT.
package property_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/property"
)

func TestOptionTable(t *testing.T) {
	o := property.NewOptionEntity()
	if o.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestOptionTableAlias(t *testing.T) {
	o := property.NewOptionEntity()
	if o.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestOptionIsUpdated(t *testing.T) {
	o := property.NewOptionEntity()
	o.IsUpdated()
}

func TestOptionID(t *testing.T) {
	o := property.NewOptionEntity()
	o.ID()
}

func TestOptionCreatedByID(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetCreatedByID(testValue)
	if testValue != o.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedByID(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetUpdatedByID(&testValue)
	if &testValue != o.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionCreatedAt(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := time.Now()
	o.SetCreatedAt(testValue)
	if testValue != o.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedAt(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := time.Now()
	o.SetUpdatedAt(&testValue)
	if &testValue != o.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionCreatedByFirstName(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetCreatedByFirstName(&testValue)
	if &testValue != o.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionCreatedBySurname(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetCreatedBySurname(&testValue)
	if &testValue != o.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedByFirstName(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetUpdatedByFirstName(&testValue)
	if &testValue != o.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionUpdatedBySurname(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetUpdatedBySurname(&testValue)
	if &testValue != o.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionActive(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := true
	o.SetActive(testValue)
	if testValue != o.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionSorting(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := uint(1e9)
	o.SetSorting(testValue)
	if testValue != o.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionKey(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetKey(&testValue)
	if &testValue != o.Key() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestOptionPropertyID(t *testing.T) {
	o := property.NewOptionEntity()
	testValue := "testValue"
	o.SetPropertyID(testValue)
	if testValue != o.PropertyID() {
		t.Fatal("Getter did not return the Set value")
	}
}
