// Code generated by espal-store-synthesizer. DO NOT EDIT.
package gift_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/gift"
)

func TestGiftTable(t *testing.T) {
	gift := gift.NewGiftEntity()
	if gift.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestGiftTableAlias(t *testing.T) {
	gift := gift.NewGiftEntity()
	if gift.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestGiftIsUpdated(t *testing.T) {
	gift := gift.NewGiftEntity()
	gift.IsUpdated()
}

func TestGiftID(t *testing.T) {
	gift := gift.NewGiftEntity()
	gift.ID()
}

func TestGiftCreatedByID(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := "testValue"
	gift.SetCreatedByID(testValue)
	if testValue != gift.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGiftUpdatedByID(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := "testValue"
	gift.SetUpdatedByID(&testValue)
	if &testValue != gift.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGiftCreatedAt(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := time.Now()
	gift.SetCreatedAt(testValue)
	if testValue != gift.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGiftUpdatedAt(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := time.Now()
	gift.SetUpdatedAt(&testValue)
	if &testValue != gift.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGiftCreatedByFirstName(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := "testValue"
	gift.SetCreatedByFirstName(&testValue)
	if &testValue != gift.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGiftCreatedBySurname(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := "testValue"
	gift.SetCreatedBySurname(&testValue)
	if &testValue != gift.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGiftUpdatedByFirstName(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := "testValue"
	gift.SetUpdatedByFirstName(&testValue)
	if &testValue != gift.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestGiftUpdatedBySurname(t *testing.T) {
	gift := gift.NewGiftEntity()
	testValue := "testValue"
	gift.SetUpdatedBySurname(&testValue)
	if &testValue != gift.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}
