// Code generated by espal-store-synthesizer. DO NOT EDIT.
package gift_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/gift"
)

func TestWrappingTable(t *testing.T) {
	w := gift.NewWrappingEntity()
	if w.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestWrappingTableAlias(t *testing.T) {
	w := gift.NewWrappingEntity()
	if w.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestWrappingIsUpdated(t *testing.T) {
	w := gift.NewWrappingEntity()
	w.IsUpdated()
}

func TestWrappingID(t *testing.T) {
	w := gift.NewWrappingEntity()
	w.ID()
}

func TestWrappingCreatedByID(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := "testValue"
	w.SetCreatedByID(testValue)
	if testValue != w.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestWrappingUpdatedByID(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := "testValue"
	w.SetUpdatedByID(&testValue)
	if &testValue != w.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestWrappingCreatedAt(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := time.Now()
	w.SetCreatedAt(testValue)
	if testValue != w.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestWrappingUpdatedAt(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := time.Now()
	w.SetUpdatedAt(&testValue)
	if &testValue != w.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestWrappingCreatedByFirstName(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := "testValue"
	w.SetCreatedByFirstName(&testValue)
	if &testValue != w.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestWrappingCreatedBySurname(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := "testValue"
	w.SetCreatedBySurname(&testValue)
	if &testValue != w.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestWrappingUpdatedByFirstName(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := "testValue"
	w.SetUpdatedByFirstName(&testValue)
	if &testValue != w.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestWrappingUpdatedBySurname(t *testing.T) {
	w := gift.NewWrappingEntity()
	testValue := "testValue"
	w.SetUpdatedBySurname(&testValue)
	if &testValue != w.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}
