// Code generated by espal-store-synthesizer. DO NOT EDIT.
package group_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/property/group"
)

func TestPropertyTable(t *testing.T) {
	p := group.NewPropertyEntity()
	if p.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestPropertyTableAlias(t *testing.T) {
	p := group.NewPropertyEntity()
	if p.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestPropertyIsUpdated(t *testing.T) {
	p := group.NewPropertyEntity()
	p.IsUpdated()
}

func TestPropertyID(t *testing.T) {
	p := group.NewPropertyEntity()
	p.ID()
}

func TestPropertyCreatedByID(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetCreatedByID(testValue)
	if testValue != p.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyUpdatedByID(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetUpdatedByID(&testValue)
	if &testValue != p.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyCreatedAt(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := time.Now()
	p.SetCreatedAt(testValue)
	if testValue != p.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyUpdatedAt(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := time.Now()
	p.SetUpdatedAt(&testValue)
	if &testValue != p.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyCreatedByFirstName(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetCreatedByFirstName(&testValue)
	if &testValue != p.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyCreatedBySurname(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetCreatedBySurname(&testValue)
	if &testValue != p.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyUpdatedByFirstName(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetUpdatedByFirstName(&testValue)
	if &testValue != p.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyUpdatedBySurname(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetUpdatedBySurname(&testValue)
	if &testValue != p.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyPropertyGroupID(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetPropertyGroupID(testValue)
	if testValue != p.PropertyGroupID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestPropertyPropertyID(t *testing.T) {
	p := group.NewPropertyEntity()
	testValue := "testValue"
	p.SetPropertyID(testValue)
	if testValue != p.PropertyID() {
		t.Fatal("Getter did not return the Set value")
	}
}
