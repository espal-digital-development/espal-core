// Code generated by espal-store-synthesizer. DO NOT EDIT.
package wishlist_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/wishlist"
)

func TestLineTable(t *testing.T) {
	l := wishlist.NewLineEntity()
	if l.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestLineTableAlias(t *testing.T) {
	l := wishlist.NewLineEntity()
	if l.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestLineIsUpdated(t *testing.T) {
	l := wishlist.NewLineEntity()
	l.IsUpdated()
}

func TestLineID(t *testing.T) {
	l := wishlist.NewLineEntity()
	l.ID()
}

func TestLineCreatedByID(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetCreatedByID(testValue)
	if testValue != l.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineUpdatedByID(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetUpdatedByID(&testValue)
	if &testValue != l.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineCreatedAt(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := time.Now()
	l.SetCreatedAt(testValue)
	if testValue != l.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineUpdatedAt(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := time.Now()
	l.SetUpdatedAt(&testValue)
	if &testValue != l.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineCreatedByFirstName(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetCreatedByFirstName(&testValue)
	if &testValue != l.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineCreatedBySurname(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetCreatedBySurname(&testValue)
	if &testValue != l.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineUpdatedByFirstName(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetUpdatedByFirstName(&testValue)
	if &testValue != l.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineUpdatedBySurname(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetUpdatedBySurname(&testValue)
	if &testValue != l.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineWishlistID(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetWishlistID(testValue)
	if testValue != l.WishlistID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineProductVariantID(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetProductVariantID(&testValue)
	if &testValue != l.ProductVariantID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineBundledProductID(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := "testValue"
	l.SetBundledProductID(&testValue)
	if &testValue != l.BundledProductID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineSorting(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := uint(1e9)
	l.SetSorting(testValue)
	if testValue != l.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestLineQuantity(t *testing.T) {
	l := wishlist.NewLineEntity()
	testValue := uint(1e9)
	l.SetQuantity(testValue)
	if testValue != l.Quantity() {
		t.Fatal("Getter did not return the Set value")
	}
}
