// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/product"
)

func TestBundledPropertyTable(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	if bundledProperty.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestBundledPropertyTableAlias(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	if bundledProperty.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestBundledPropertyIsUpdated(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	bundledProperty.IsUpdated()
}

func TestBundledPropertyID(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	bundledProperty.ID()
}

func TestBundledPropertyCreatedByID(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetCreatedByID(testValue)
	if testValue != bundledProperty.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyUpdatedByID(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetUpdatedByID(&testValue)
	if &testValue != bundledProperty.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyCreatedAt(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := time.Now()
	bundledProperty.SetCreatedAt(testValue)
	if testValue != bundledProperty.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyUpdatedAt(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := time.Now()
	bundledProperty.SetUpdatedAt(&testValue)
	if &testValue != bundledProperty.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyCreatedByFirstName(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetCreatedByFirstName(&testValue)
	if &testValue != bundledProperty.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyCreatedBySurname(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetCreatedBySurname(&testValue)
	if &testValue != bundledProperty.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyUpdatedByFirstName(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetUpdatedByFirstName(&testValue)
	if &testValue != bundledProperty.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyUpdatedBySurname(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetUpdatedBySurname(&testValue)
	if &testValue != bundledProperty.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertySorting(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := uint(1e9)
	bundledProperty.SetSorting(testValue)
	if testValue != bundledProperty.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyKey(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetKey(&testValue)
	if &testValue != bundledProperty.Key() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyBundledProductID(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetBundledProductID(testValue)
	if testValue != bundledProperty.BundledProductID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestBundledPropertyPropertyID(t *testing.T) {
	bundledProperty := product.NewBundledPropertyEntity()
	testValue := "testValue"
	bundledProperty.SetPropertyID(testValue)
	if testValue != bundledProperty.PropertyID() {
		t.Fatal("Getter did not return the Set value")
	}
}
