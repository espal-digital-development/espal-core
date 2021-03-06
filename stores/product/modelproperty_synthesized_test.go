// Code generated by espal-store-synthesizer. DO NOT EDIT.
package product_test

import (
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/stores/product"
)

func TestModelPropertyTable(t *testing.T) {
	m := product.NewModelPropertyEntity()
	if m.TableName() == "" {
		t.Fatal("TableName shouldn't be empty")
	}
}

func TestModelPropertyTableAlias(t *testing.T) {
	m := product.NewModelPropertyEntity()
	if m.TableName() == "" {
		t.Fatal("TableAlias shouldn't be empty")
	}
}

func TestModelPropertyIsUpdated(t *testing.T) {
	m := product.NewModelPropertyEntity()
	m.IsUpdated()
}

func TestModelPropertyID(t *testing.T) {
	m := product.NewModelPropertyEntity()
	m.ID()
}

func TestModelPropertyCreatedByID(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetCreatedByID(testValue)
	if testValue != m.CreatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyUpdatedByID(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetUpdatedByID(&testValue)
	if &testValue != m.UpdatedByID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyCreatedAt(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := time.Now()
	m.SetCreatedAt(testValue)
	if testValue != m.CreatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyUpdatedAt(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := time.Now()
	m.SetUpdatedAt(&testValue)
	if &testValue != m.UpdatedAt() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyCreatedByFirstName(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetCreatedByFirstName(&testValue)
	if &testValue != m.CreatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyCreatedBySurname(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetCreatedBySurname(&testValue)
	if &testValue != m.CreatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyUpdatedByFirstName(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetUpdatedByFirstName(&testValue)
	if &testValue != m.UpdatedByFirstName() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyUpdatedBySurname(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetUpdatedBySurname(&testValue)
	if &testValue != m.UpdatedBySurname() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyProductModelID(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetProductModelID(testValue)
	if testValue != m.ProductModelID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyPropertyID(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetPropertyID(testValue)
	if testValue != m.PropertyID() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertySorting(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := uint(1e9)
	m.SetSorting(testValue)
	if testValue != m.Sorting() {
		t.Fatal("Getter did not return the Set value")
	}
}

func TestModelPropertyKey(t *testing.T) {
	m := product.NewModelPropertyEntity()
	testValue := "testValue"
	m.SetKey(&testValue)
	if &testValue != m.Key() {
		t.Fatal("Getter did not return the Set value")
	}
}
