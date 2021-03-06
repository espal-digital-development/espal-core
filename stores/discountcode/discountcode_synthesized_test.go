// Code generated by espal-store-synthesizer. DO NOT EDIT.
package discountcode_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/discountcode"
)

func TestDiscountCodeTable(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	if d.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestDiscountCodeTableAlias(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	if d.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestDiscountCodeIsUpdated(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	d.IsUpdated()
}

func TestDiscountCodeID(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	d.ID()
}

func TestDiscountCodeCreatedByID(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := "testValue"
	d.SetCreatedByID(testValue)
	if testValue != d.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeUpdatedByID(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := "testValue"
	d.SetUpdatedByID(&testValue)
	if &testValue != d.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeCreatedAt(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := time.Now()
	d.SetCreatedAt(testValue)
	if testValue != d.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeUpdatedAt(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := time.Now()
	d.SetUpdatedAt(&testValue)
	if &testValue != d.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeCreatedByFirstName(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := "testValue"
	d.SetCreatedByFirstName(&testValue)
	if &testValue != d.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeCreatedBySurname(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := "testValue"
	d.SetCreatedBySurname(&testValue)
	if &testValue != d.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeUpdatedByFirstName(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := "testValue"
	d.SetUpdatedByFirstName(&testValue)
	if &testValue != d.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeUpdatedBySurname(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := "testValue"
	d.SetUpdatedBySurname(&testValue)
	if &testValue != d.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeKey(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := "testValue"
	d.SetKey(testValue)
	if testValue != d.Key() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeMaxUses(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := uint(1e9)
	d.SetMaxUses(&testValue)
	if &testValue != d.MaxUses() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeUsesCounter(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := uint(1e9)
	d.SetUsesCounter(testValue)
	if testValue != d.UsesCounter() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeAvailableFrom(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := time.Now()
	d.SetAvailableFrom(&testValue)
	if &testValue != d.AvailableFrom() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeAvailableUntil(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := time.Now()
	d.SetAvailableUntil(&testValue)
	if &testValue != d.AvailableUntil() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeDiscountPercentage(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := float32(3.14)
	d.SetDiscountPercentage(&testValue)
	if &testValue != d.DiscountPercentage() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestDiscountCodeDiscountAmount(t *testing.T) {
	d := discountcode.NewDiscountCodeEntity()
	testValue := float32(3.14)
	d.SetDiscountAmount(&testValue)
	if &testValue != d.DiscountAmount() {
		t.Fatal("Getter did not return the Set value")
	}
}
