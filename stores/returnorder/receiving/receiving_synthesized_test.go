// Code generated by espal-store-synthesizer. DO NOT EDIT.
package receiving_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/returnorder/receiving"
)

func TestReceivingTable(t *testing.T) {
	r := receiving.NewReceivingEntity()
	if r.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestReceivingTableAlias(t *testing.T) {
	r := receiving.NewReceivingEntity()
	if r.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestReceivingIsUpdated(t *testing.T) {
	r := receiving.NewReceivingEntity()
	r.IsUpdated()
}

func TestReceivingID(t *testing.T) {
	r := receiving.NewReceivingEntity()
	r.ID()
}

func TestReceivingCreatedByID(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetCreatedByID(testValue)
	if testValue != r.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingUpdatedByID(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetUpdatedByID(&testValue)
	if &testValue != r.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingCreatedAt(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := time.Now()
	r.SetCreatedAt(testValue)
	if testValue != r.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingUpdatedAt(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := time.Now()
	r.SetUpdatedAt(&testValue)
	if &testValue != r.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingCreatedByFirstName(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetCreatedByFirstName(&testValue)
	if &testValue != r.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingCreatedBySurname(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetCreatedBySurname(&testValue)
	if &testValue != r.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingUpdatedByFirstName(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetUpdatedByFirstName(&testValue)
	if &testValue != r.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingUpdatedBySurname(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetUpdatedBySurname(&testValue)
	if &testValue != r.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingDomainID(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetDomainID(testValue)
	if testValue != r.DomainID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestReceivingComments(t *testing.T) {
	r := receiving.NewReceivingEntity()
	testValue := "testValue"
	r.SetComments(&testValue)
	if &testValue != r.Comments() {
		t.Fatal("Getter did not return the Set value")
	}
}
