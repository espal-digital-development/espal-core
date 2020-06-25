// Code generated by espal-store-synthesizer. DO NOT EDIT.
package reseller_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/reseller"
)

func TestResellerTable(t *testing.T) {
	r := reseller.NewResellerEntity()
	if r.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestResellerTableAlias(t *testing.T) {
	r := reseller.NewResellerEntity()
	if r.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestResellerIsUpdated(t *testing.T) {
	r := reseller.NewResellerEntity()
	r.IsUpdated()
}

func TestResellerID(t *testing.T) {
	r := reseller.NewResellerEntity()
	r.ID()
}

func TestResellerCreatedByID(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := "testValue"
	r.SetCreatedByID(testValue)
	if testValue != r.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerUpdatedByID(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := "testValue"
	r.SetUpdatedByID(&testValue)
	if &testValue != r.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerCreatedAt(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := time.Now()
	r.SetCreatedAt(testValue)
	if testValue != r.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerUpdatedAt(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := time.Now()
	r.SetUpdatedAt(&testValue)
	if &testValue != r.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerCreatedByFirstName(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := "testValue"
	r.SetCreatedByFirstName(&testValue)
	if &testValue != r.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerCreatedBySurname(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := "testValue"
	r.SetCreatedBySurname(&testValue)
	if &testValue != r.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerUpdatedByFirstName(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := "testValue"
	r.SetUpdatedByFirstName(&testValue)
	if &testValue != r.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerUpdatedBySurname(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := "testValue"
	r.SetUpdatedBySurname(&testValue)
	if &testValue != r.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerActive(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := true
	r.SetActive(testValue)
	if testValue != r.Active() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerCountry(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := uint16(65000)
	r.SetCountry(&testValue)
	if &testValue != r.Country() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestResellerAddress(t *testing.T) {
	r := reseller.NewResellerEntity()
	testValue := "testValue"
	r.SetAddress(&testValue)
	if &testValue != r.Address() {
		t.Fatal("Getter did not return the Set value")
	}
}
